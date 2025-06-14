package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
)

const (
	creditsPrefix       = "credits_"
	HARDCODED_X_API_KEY = "Bearer twjvq4Cvs8NJDNmuCxTruE3QkX3DLwVPNYjPPu8AFSFj2ELnPJumJTM7c1pV"
)

// ConfirmTransaction godoc
// @Summary Confirm an IAP transaction from stores
// @Description Confirm an IAP transaction from stores
// @ID CreateRevenueCatPurchase
// @Produce json
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /rc/create-transaction [post]
func (e *Env) CreateRevenueCatPurchase(c echo.Context) error {
	var req model.RevenueCatPurchaseRequest
	apiKey := c.Request().Header.Get("Authorization")
	if apiKey != HARDCODED_X_API_KEY {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	// fmt.Printf("\n%+v\n", req.Event)
	if req.Event == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing required fields")
	}
	if req.Event.ProductID == "" || req.Event.Store == "" || req.Event.Type == "" || req.Event.AppUserID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing required fields")
	}
	// check if the user exists
	if _, err := e.Store.GetUserByID(req.Event.AppUserID); err != nil {
		log.Printf("CreateRevenueCatPurchase - retrieving user [%s] err: %v", req.Event.AppUserID, err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	// check if the product exists
	product, err := e.Store.GetStoreProductByStoreAndProductID(req.Event.Store, req.Event.ProductID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid product")
	}
	// check if the product is active
	if !product.Active {
		return echo.NewHTTPError(http.StatusBadRequest, "Product not active")
	}
	// convert req.Event.PurchasedAtMs from unix epoch to time.Time
	purchasedAt := time.Unix(req.Event.PurchasedAtMs/1000, 0)

	// create the store product transaction
	stID, err := e.Store.CreateStoreProductTransaction(
		req.Event.OriginalTransactionID,
		req.Event.Store,
		product.ID,
		purchasedAt,
	)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	transactionObjectType := database.TransactionTypeVirtual
	if product.ProductType.String == database.StoreProductTypeNFT {
		transactionObjectType = database.TransactionTypeNFT
	}
	// helper function to parse actions to be made for the just created transaction
	if err := e.parseTransactionAction(
		stID, req.Event.OriginalTransactionID, req.Event.AppUserID,
		product.StoreProductID, transactionObjectType,
	); err != nil {
		return err
	}

	return e.RespondSuccess(c, "ok")
}

func (e *Env) parseTransactionAction(storeTransactionID, externalTransactionID, userID, productCode, transactionObjectType string) *echo.HTTPError {
	if userID == "" {
		log.Println("User ID is missing")
		return echo.NewHTTPError(http.StatusBadRequest, "User ID is missing")
	}
	description := fmt.Sprintf("Purchase of [%s] product with transaction ID: %s", productCode, externalTransactionID)
	amount := 0.0

	// Find store product by product id
	storeProduct, err := e.Store.GetStoreProductByProductID(productCode)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !storeProduct.Active {
		return echo.NewHTTPError(http.StatusBadRequest, "Product is not active")
	}

	// Pre-confirm purchase actions
	// -------------------------------------------------------------------------
	// Credits action
	if strings.HasPrefix(productCode, creditsPrefix) {
		amountStr := strings.TrimPrefix(productCode, creditsPrefix)
		if parsedAmount, err := strconv.ParseFloat(amountStr, 64); err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid amount")
		} else {
			amount = parsedAmount
		}
	}
	// -------------------------------------------------------------------------

	trx, err := e.Store.CreateTransaction(
		&schema.Transaction{
			ID:         uuid.NewString(),
			UserID:     userID,
			Amount:     amount,
			ObjectType: transactionObjectType,
			Quantity:   int64(1),
			Text:       null.StringFrom(description),
		},
	)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else {
		// confirm the store product transaction
		if err := e.Store.ConfirmStoreProductTransactionByExternalID(
			userID,
			trx.ID,
			externalTransactionID,
		); err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	// Post-confirm purchase actions
	// -------------------------------------------------------------------------
	// Card pack action - if the store product is a cardpack, assign it to the user
	var assignedCardPack *schema.AssignedCardPack
	if cardPackType, err := e.Store.GetCardPackTypeByCode(storeProduct.StoreProductID); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	} else {
		// assign card pack
		assignedCardPack = &schema.AssignedCardPack{
			ID:                 uuid.New().String(),
			UserID:             userID,
			StoreTransactionID: null.StringFrom(storeTransactionID),
			CardPackTypeID:     cardPackType.ID,
			Opened:             false,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}
		err := e.Store.DeductFromPackLimit(cardPackType.ID, 1)
		if err != nil {
			log.Printf("IAP - Error deducting from pack limit: %v", err)
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				map[string]string{"error": fmt.Sprintf("Failed to assign card pack: %s", err.Error())},
			)
		}
		if _, err := e.Store.CreateAssignedCardPack(assignedCardPack); err != nil {
			log.Printf("IAP - Error creating assigned card pack: %v", err)
			errMap := map[string]string{"error": fmt.Sprintf("IAP - Failed to assign card pack: %s", err.Error())}
			if err := e.Store.RestockPackLimit(cardPackType.ID, 1); err != nil {
				errMap["error"] = errMap["error"] + " - RestockPackLimit: " + err.Error()
			}
			return echo.NewHTTPError(http.StatusInternalServerError, errMap)
		}
	}
	// -------------------------------------------------------------------------
	if err := e.processReferralBonus(userID, trx, assignedCardPack); err != nil {
		logrus.WithError(err).Error("cannot process referral bonus")
	}

	return nil
}

func (e *Env) processReferralBonus(
	userID string,
	originalTrx *schema.Transaction,
	assignedCardPack *schema.AssignedCardPack,
) error {
	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("cannot get user: %w", err)
	}

	// referral bonus already used
	if user.ReferralBonusUsed.Bool {
		return nil
	}

	// user don't have a referrer, too bad for him, set like he used a bonus
	if !user.ReferrerID.Valid {
		user.ReferralBonusUsed = null.BoolFrom(true)
		return e.Store.UpdateUserReferrerBonusUsed(user)
	}

	return e.Store.Transaction(
		func(s database.Store) error {
			// make credits reward
			_, err := e.Store.CreateTransaction(
				&schema.Transaction{
					ID:         uuid.New().String(),
					UserID:     user.ReferrerID.String,
					Amount:     originalTrx.Amount,
					ObjectType: originalTrx.ObjectType,
					Text:       null.StringFrom("Referral gift"),
					Quantity:   originalTrx.Quantity,
				},
			)
			if err != nil {
				return fmt.Errorf("cannot create transaction: %w", err)
			}

			if assignedCardPack != nil {
				// made card pack reward
				err := e.Store.DeductFromPackLimit(assignedCardPack.CardPackTypeID, 1)
				if err != nil {
					return fmt.Errorf("cannot deduct from pack limit: %w", err)
				}

				newAssignedCardPack := &schema.AssignedCardPack{
					ID:                 uuid.New().String(),
					UserID:             user.ReferrerID.String,
					StoreTransactionID: null.StringFrom(assignedCardPack.StoreTransactionID.String),
					CardPackTypeID:     assignedCardPack.CardPackTypeID,
					Opened:             false,
					CreatedAt:          time.Now(),
					UpdatedAt:          time.Now(),
				}
				if _, err := e.Store.CreateAssignedCardPack(newAssignedCardPack); err != nil {
					return fmt.Errorf("cannot create assigned card pack: %w", err)
				}
			}

			logrus.WithField("user_id", user.ID).
				WithField("referrer_id", user.ReferrerID).
				Info("created referral reward for referrer user")

			user.ReferralBonusUsed = null.BoolFrom(true)
			return s.UpdateUserReferrerBonusUsed(user)
		},
	)
}

func (e *Env) GetCardPackTypes(c echo.Context) error {
	cardPackTypes, err := e.Store.GetAllCardPacks()
	if err != nil {
		return err
	}
	parsedPacks := make([]model.CardPackType, 0)
	for _, cardPackType := range cardPackTypes {
		var packLimits int64
		if cardPackType.PackLimits.Valid {
			packLimits = int64(cardPackType.PackLimits.Int)
		}
		parsedPacks = append(
			parsedPacks, model.CardPackType{
				ID:           cardPackType.ID,
				Name:         cardPackType.Name,
				PackLimits:   packLimits,
				CardPackCode: cardPackType.CardPackCode.String,
			},
		)
	}
	return e.RespondSuccess(c, parsedPacks)
}

// CancelRevenueCatPurchase godoc
// @Summary Cancels an IAP transaction from stores
// @Description Cancels an IAP transaction from stores
// @ID CancelRevenueCatPurchase
// @Produce json
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /rc/create-transaction [post]
func (e *Env) CancelRevenueCatPurchase(c echo.Context) error {
	var req model.RevenueCatPurchaseRequest
	apiKey := c.Request().Header.Get("Authorization")
	if apiKey != HARDCODED_X_API_KEY {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	// fmt.Printf("\n%+v\n", req.Event)
	if req.Event == nil && (req.Event.ProductID == "" || req.Event.Store == "" || req.Event.Type == "" || req.Event.AppUserID == "") {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing required fields")
	}
	if req.Event.Type != "CANCELLATION" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid event request type")
	}
	// check if the user exists
	if _, err := e.Store.GetUserByID(req.Event.AppUserID); err != nil {
		log.Printf("CreateRevenueCatPurchase - retrieving user [%s] err: %v", req.Event.AppUserID, err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	// check if the product exists
	product, err := e.Store.GetStoreProductByStoreAndProductID(req.Event.Store, req.Event.ProductID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid product")
	}
	// check if the product is active
	if !product.Active {
		log.Println("Product not active - ", req.Event.ProductID)
		return echo.NewHTTPError(http.StatusBadRequest, "Product not active")
	}

	// create the store product transaction
	spTrx, err := e.Store.GetStoreProductTransactionByExternalID(req.Event.OriginalTransactionID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// get the assginedcard pack by store transaction id
	acp, err := e.Store.GetAssignedCardPackByStoreProductTransactionID(spTrx.ID)
	if err != nil {
		log.Println("[GetAssignedCardPackByStoreProductTransactionID] ERROR " + spTrx.ID + " - " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// refund the assiged_card_pack
	if err := e.Store.RefundAssignedCardPackByID(acp.ID); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.RespondSuccess(c, "ok")
}
