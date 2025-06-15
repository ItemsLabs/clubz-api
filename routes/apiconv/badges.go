package apiconv

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
)

// ToBadgeSlice converts a slice of Badge schema objects to a slice of Badge models.
func ToBadgeSlice(arr schema.BadgeSlice) []*model.Badge {
	result := make([]*model.Badge, 0, len(arr))
	for _, el := range arr {
		result = append(result, ToBadge(el))
	}

	return result
}

// ToBadge converts a Badge schema object to a Badge model.
func ToBadge(el *schema.Badge) *model.Badge {
	return &model.Badge{
		ID:          int64(el.ID),
		Name:        el.Name,
		Description: el.Description,
		Image:       el.Image,
		Points:      int64(el.Points),
		Type:        el.Type,
		Status:      el.Status,
	}
}
