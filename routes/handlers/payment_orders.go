package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/itemslabs/clubz-api/config"
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/apiconv"
	"github.com/itemslabs/clubz-api/routes/handlers/payments"
	"github.com/itemslabs/clubz-api/types"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v76/webhook"
	"github.com/volatiletech/null/v8"
)

type Provider interface {
	Purchase(user *schema.User, purchaseDetails types.PurchaseDetails) (*schema.Order, error)
}

// Purchase godoc
// @Summary Purchase endpoint for payments providers
// @Description Purchase endpoint for payments providers
// @ID purchase
// @Produce json
// @Param body PurchaseRequest true "Purchase request"
// @Success 200 {object} Order
// @Router /purchase [post]

func (e *Env) Purchase(c echo.Context) error {
	var purchaseDetails types.PurchaseDetails
	// parse body
	if err := e.ParseBody(c, &purchaseDetails); err != nil {
		return err
	}
	// check user
	user, err := e.Store.GetUserByID(userID(c))
	if err != nil {
		return err
	}

	var paymentProvider Provider

	STRIPE := int64(1)
	INTERNAL := int64(2)

	switch purchaseDetails.PaymentMethod {
	case STRIPE:
		paymentProvider = &payments.StripePaymentProvider{SecretKey: config.StripeSecret(), Store: e.Store}
	case INTERNAL:
		paymentProvider = &payments.InternalPaymentProvider{Store: e.Store}
	default:
		return ErrPaymentMethod
	}

	order, err := paymentProvider.Purchase(user, purchaseDetails)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error with Order: %s", err))
	}

	return e.RespondSuccess(c, apiconv.ToOrder(order))
}

func (e *Env) WebhooksPurchase(c echo.Context) error {
	req := c.Request()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error reading request body")
	}
	defer req.Body.Close()

	event, err := webhook.ConstructEventWithOptions(body, c.Request().Header.Get("Stripe-Signature"), config.StripeSigningSecret(), webhook.ConstructEventOptions{IgnoreAPIVersionMismatch: true})
	if err != nil {
		return c.String(http.StatusBadRequest, "Error reading event")
	}
	orderID := event.GetObjectValue("client_reference_id")
	order, err := e.Store.GetOrderByID(orderID)
	if err != nil {
		return c.String(http.StatusNotFound, "Order not found")
	}
	CompletePaymentStatus := 2

	if order.PaymentPlatformStatus == CompletePaymentStatus {
		return c.String(http.StatusBadRequest, "Order already completed")
	}

	if order.Delivered {
		return c.String(http.StatusBadRequest, "Order already delivered")
	}

	ProcessingBlockchainStatus := 1
	switch event.Type {
	case "checkout.session.completed":
		println("checkout.session.completed")
		order.PurchasedAt = null.TimeFrom(time.Now())
		order.PaymentPlatformStatus = CompletePaymentStatus
		order.BlockchainOrderStatus = ProcessingBlockchainStatus
		log.Println("minting imitation")

		order.BlockchainUUID = null.StringFrom("retrieved transactionID from blockchain")

		err = e.Store.UpdateOrder(order)
		if err != nil {
			return c.String(http.StatusBadRequest, "Error update Order")
		}

	case "checkout.session.expired":
		order.PaymentPlatformStatus = 3 // expired
		err = e.Store.UpdateOrder(order)
		if err != nil {
			return c.String(http.StatusBadRequest, "Error update Order")
		}
	default:
		return c.String(http.StatusBadRequest, fmt.Sprintf("Handler for event type: %s missing", event.Type))
	}

	return e.RespondSuccess(c, nil)
}

// GetOrderByID godoc
// @Summary Get Order by ID
// @Description Get Order by ID
// @ID get-order-by-id
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} Order
// @Router /purchase/{id} [get]

func (e *Env) GetOrderByID(c echo.Context) error {
	order, err := e.Store.GetOrderByID(c.Param("id"))

	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToOrder(order))
}
