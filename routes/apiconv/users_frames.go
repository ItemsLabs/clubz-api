package apiconv

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/go-openapi/strfmt"
)

// ToUserFrameSlice converts a slice of UserFrame schema objects to a slice of UserFrame models.
func ToUserFrameSlice(arr schema.UserFrameSlice) []*model.UserFrame {
	result := make([]*model.UserFrame, 0, len(arr))
	for _, el := range arr {
		result = append(result, ToUserFrame(el))
	}

	return result
}

// ToUserFrame converts a UserFrame schema object to a UserFrame model.
func ToUserFrame(el *schema.UserFrame) *model.UserFrame {
	return &model.UserFrame{
		ID:        int64(el.ID),
		UserID:    el.UserID,
		FrameID:   int64(el.FrameID),
		CreatedAt: strfmt.DateTime(el.CreatedAt),
		UpdatedAt: strfmt.DateTime(el.UpdatedAt),
	}
}
