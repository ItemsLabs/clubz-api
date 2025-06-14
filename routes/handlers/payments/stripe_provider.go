package payments

import (
	"strconv"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/types"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"github.com/volatiletech/null/v8"
)

type StripePaymentProvider struct {
	SecretKey string
	Store     database.Store
}

type Metadata struct {
	TrackerId             int64  `json:"tracker_id"`
	Quantity              int64  `json:"quantity"`
	BlockchainOrderStatus int64  `json:"blockchain_order_status"`
	WalletAddress         string `json:"wallet_address"`
	BlockchainUuid        string `json:"blockchain_uuid"`
	Delivered             bool   `json:"delivered"`
}

type MetadataMap map[string]string

func (s *StripePaymentProvider) Purchase(user *schema.User, purchaseDetails types.PurchaseDetails) (*schema.Order, error) {
	item, err := s.Store.GetItemByID(purchaseDetails.ItemId)
	if err != nil {
		return nil, err
	}

	PaymentPlatformType := 1

	var orderParams = &schema.Order{
		ID:                  uuid.New().String(),
		Quantity:            purchaseDetails.Quantity,
		Amount:              purchaseDetails.Amount,
		Contract:            null.StringFrom(item.ContractAddress),
		TokenID:             null.StringFrom(item.TokenID.String),
		Delivered:           false,
		PaymentPlatformType: PaymentPlatformType,
		CancelURL:           null.StringFrom(purchaseDetails.CancelUrl),
		SuccessURL:          null.StringFrom(purchaseDetails.SuccessUrl),
		ItemID:              null.IntFrom(item.ID),
		UserID:              user.ID,
	}
	order, err := s.Store.CreateOrder(orderParams)
	if err != nil {
		return nil, err
	}

	successPageUrl := purchaseDetails.SuccessUrl + "?order_id=" + order.ID

	paymentSession, err := s.CreatePaymentIntent(purchaseDetails.Quantity, order.ID, item.StripePriceID.String, successPageUrl, purchaseDetails.CancelUrl)
	if err != nil {
		return nil, err
	}

	order.PaymentPlatformUUID = null.StringFrom(paymentSession.ID)
	order.PaymentPlatformURL = null.StringFrom(paymentSession.URL)
	err = s.Store.UpdateOrder(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *StripePaymentProvider) CreatePaymentIntent(quantity int64, trackerId, priceID, successUrl, cancelUrl string) (*stripe.CheckoutSession, error) {
	stripe.Key = s.SecretKey
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceID),
				Quantity: stripe.Int64(quantity),
			},
		},
		Mode:                     stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:               stripe.String(successUrl),
		CancelURL:                stripe.String(cancelUrl),
		ClientReferenceID:        stripe.String(trackerId),
		BillingAddressCollection: stripe.String("required"),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{
			Enabled: stripe.Bool(true),
		},
	}

	stripeSession, err := session.New(params)
	if err != nil {
		return nil, err
	}

	return stripeSession, nil
}

func (*StripePaymentProvider) PushMetadata(paymentIntentId string, metadata Metadata) error {

	params := &stripe.PaymentIntentParams{}
	meta := MetadataMap{
		"purchase_tracker_id":     strconv.FormatInt(metadata.TrackerId, 10),
		"buyer_wallet_address":    metadata.WalletAddress,
		"quantity":                strconv.FormatInt(metadata.Quantity, 10),
		"blockchain_uuid":         metadata.BlockchainUuid,
		"blockchain_order_status": strconv.FormatInt(metadata.BlockchainOrderStatus, 10),
		"delivered":               strconv.FormatBool(metadata.Delivered),
	}

	params.Metadata = meta
	_, err := paymentintent.Update(paymentIntentId, params)
	return err
}
