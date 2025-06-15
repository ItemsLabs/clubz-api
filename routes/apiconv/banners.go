package apiconv

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
)

// ToBannerSlice converts a slice of Banner schema objects to a slice of Banner models.
func ToBannerSlice(arr schema.BannerSlice) []*model.Banner {
	result := make([]*model.Banner, 0, len(arr))
	for _, el := range arr {
		result = append(result, ToBanner(el))
	}

	return result
}

// ToBanner converts a Banner schema object to a Banner model.
func ToBanner(el *schema.Banner) *model.Banner {
	return &model.Banner{
		ID:          int64(el.ID),
		Name:        el.Name,
		Description: el.Description,
		Image:       el.Image,
		Points:      int64(el.Points),
		Type:        el.Type,
		Status:      el.Status,
	}
}
