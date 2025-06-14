package apiconv

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/go-openapi/strfmt"
)

func ToItem(item *schema.Item) *model.Item {
	return &model.Item{
		ID:                int64(item.ID),
		Price:             item.Price,
		Title:             item.Title.String,
		Description:       item.Description.String,
		PageURL:           item.PageURL.String,
		PurchaseImgURL:    item.PurchaseImgURL.String,
		ContractAbbr:      item.ContractAbbr,
		ContractAddress:   item.ContractAddress,
		TokenID:           item.TokenID.String,
		StripePriceID:     item.StripePriceID.String,
		MinQuantity:       int64(item.MinQuantity.Int),
		DefaultQuantity:   int64(item.DefaultQuantity.Int),
		MaxQuantity:       int64(item.MaxQuantity.Int),
		WhitelistRequired: item.WhitelistRequired,
		StartDateAt:       strfmt.DateTime(item.StartDateAt),
		CloseDateAt:       strfmt.DateTime(item.CloseDateAt),
		BonusQuantity:     item.BonusQuantity,
		Type:              int64(item.Type),
	}
}

func ToItemSlice(arr schema.ItemSlice) []*model.Item {
	result := make([]*model.Item, 0, len(arr))
	for _, el := range arr {
		result = append(result, ToItem(el))
	}

	return result
}
