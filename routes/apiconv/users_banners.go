package apiconv

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/go-openapi/strfmt"
)

// ToUserBannerSlice converts a slice of UserBanner schema objects to a slice of UserBanner models.
func ToUserBannerSlice(arr schema.UserBannerSlice) []*model.UserBanner {
	result := make([]*model.UserBanner, 0, len(arr))
	for _, el := range arr {
		result = append(result, ToUserBanner(el))
	}

	return result
}

// ToUserBanner converts a UserBanner schema object to a UserBanner model.
func ToUserBanner(el *schema.UserBanner) *model.UserBanner {
	return &model.UserBanner{
		ID:        int64(el.ID),
		UserID:    el.UserID,
		BannerID:  int64(el.BannerID),
		CreatedAt: strfmt.DateTime(el.CreatedAt),
		UpdatedAt: strfmt.DateTime(el.UpdatedAt),
	}
}
