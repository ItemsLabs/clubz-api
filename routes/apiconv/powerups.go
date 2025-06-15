package apiconv

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
)

func ToPowerUpSlice(arr schema.PowerupSlice, actions map[int]schema.ActionSlice) []*model.PowerUp {
	result := make([]*model.PowerUp, 0, len(arr))
	for _, el := range arr {
		var puActions schema.ActionSlice
		if v, ok := actions[el.ID]; ok {
			puActions = v
		}

		result = append(result, ToPowerUp(el, puActions))
	}

	return result
}

func ToPowerUp(pu *schema.Powerup, actions schema.ActionSlice) *model.PowerUp {
	return &model.PowerUp{
		ID:          int64(pu.ID),
		Name:        pu.Name,
		Description: pu.Description.String,
		Duration:    int64(pu.Duration),
		IconURL:     pu.IconURL.Ptr(),
		Multiplier:  ToFloatWithZero(pu.Multiplier),
		Actions:     ToActionSlice(actions),
		Cost:        ToFloatWithZero(pu.Cost),
	}
}
