package apiconv

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/go-openapi/strfmt"
)

// ToUserBadgeSlice converts a slice of UserBadge schema objects to a slice of UserBadge models.
func ToUserBadgeSlice(arr schema.UserBadgeSlice) []*model.UserBadge {
	result := make([]*model.UserBadge, 0, len(arr))
	for _, el := range arr {
		result = append(result, ToUserBadge(el))
	}

	return result
}

// ToUserBadge converts a UserBadge schema object to a UserBadge model.
func ToUserBadge(el *schema.UserBadge) *model.UserBadge {
	return &model.UserBadge{
		ID:        int64(el.ID),
		UserID:    el.UserID,
		BadgeID:   int64(el.BadgeID),
		CreatedAt: strfmt.DateTime(el.CreatedAt),
		UpdatedAt: strfmt.DateTime(el.UpdatedAt),
	}
}
