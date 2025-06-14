package apiconv

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
)

func ToActionSlice(arr schema.ActionSlice) []*model.Action {
	result := make([]*model.Action, 0, len(arr))
	for _, el := range arr {
		result = append(result, ToAction(el))
	}

	return result
}

func ToAction(el *schema.Action) *model.Action {
	return &model.Action{
		ID:          int64(el.ID),
		Name:        el.Name,
		Description: el.Description,
		Score:       ToFloatWithZero(el.Score),
		Category:    el.Category.String,
		Icon:        el.Icon.String,
	}
}
