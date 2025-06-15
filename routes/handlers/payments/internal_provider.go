package payments

import (
	"fmt"
	"log"
	"time"

	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/types"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
)

type InternalPaymentProvider struct {
	Store database.Store
}

func (i InternalPaymentProvider) Purchase(user *schema.User, purchaseDetails types.PurchaseDetails) (*schema.Order, error) {
	item, err := i.Store.GetItemByID(purchaseDetails.ItemId)
	if err != nil {
		return nil, err
	}
	// only for credits
	if item.Type == 2 {
		return nil, fmt.Errorf("you can only buy credits for GAME currency")
	}
	// TODO: implement balance checker
	//balance,err := checkWalletBalance(user)
	//if balance < (item.Price * purchaseDetails.Quantity){
	//	return nil, fmt.Errorf("you don't have enough GAME currency for buy credits")
	//}
	InternalPaymentType := 2
	CompletePaymentStatus := 2
	ProcessingBlockchainStatus := 1

	var orderParams = &schema.Order{
		ID:                    uuid.New().String(),
		Quantity:              purchaseDetails.Quantity,
		Amount:                purchaseDetails.Amount,
		Contract:              null.StringFrom(item.ContractAddress),
		TokenID:               null.StringFrom(item.TokenID.String),
		Delivered:             false,
		PaymentPlatformType:   InternalPaymentType,
		ItemID:                null.IntFrom(item.ID),
		UserID:                user.ID,
		PurchasedAt:           null.TimeFrom(time.Now()),
		PaymentPlatformStatus: CompletePaymentStatus,
	}

	order, err := i.Store.CreateOrder(orderParams)
	if err != nil {
		return nil, err
	}
	// TODO: Put Code for transferGameToken
	// transferGameToken, err := i.transferGameToken(purchaseDetails.Quantity, order.ID, item.StripePriceID.String, user.WalletAddress)
	//	if err != nil {
	//		return nil, err
	//	}

	log.Println("transferGameToken imitation")

	orderParams.BlockchainUUID = null.StringFrom("0x6fe0afd6ae9bd97471f8a24645b43f3857e995d6b4be95d10813dc523575fa03")
	order.BlockchainOrderStatus = ProcessingBlockchainStatus

	err = i.Store.UpdateOrder(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
