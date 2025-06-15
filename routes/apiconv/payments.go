package apiconv

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/go-openapi/strfmt"
)

func ToOrder(order *schema.Order) *model.Order {
	return &model.Order{
		ID:                    order.ID,
		CreatedAt:             strfmt.DateTime(order.CreatedAt),
		UpdatedAt:             strfmt.DateTime(order.UpdatedAt),
		Quantity:              int32(order.Quantity),
		Contract:              order.Contract.String,
		TokenID:               order.TokenID.String,
		Delivered:             order.Delivered,
		PurchasedAt:           strfmt.DateTime(order.PurchasedAt.Time),
		BlockchainUUID:        order.BlockchainUUID.String,
		PaymentPlatformUUID:   order.PaymentPlatformUUID.String,
		BlockchainOrderStatus: model.BlockchainStatus(order.BlockchainOrderStatus),
		PaymentPlatformStatus: model.PaymentPlatformStatus(order.PaymentPlatformStatus),
		PaymentPlatformType:   model.PaymentType(order.PaymentPlatformType),
		PaymentPlatformURL:    order.PaymentPlatformURL.String,
		CancelURL:             order.CancelURL.String,
		SuccessURL:            order.SuccessURL.String,
		ItemID:                int32(order.ItemID.Int),
		UserID:                order.UserID,
		Amount:                order.Amount,
	}
}
