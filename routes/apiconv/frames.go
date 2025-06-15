package apiconv

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
)

// ToFrameSlice converts a slice of Frame schema objects to a slice of Frame models.
func ToFrameSlice(arr schema.FrameSlice) []*model.Frame {
	result := make([]*model.Frame, 0, len(arr))
	for _, el := range arr {
		result = append(result, ToFrame(el))
	}

	return result
}

// ToFrame converts a Frame schema object to a Frame model.
func ToFrame(el *schema.Frame) *model.Frame {
	return &model.Frame{
		ID:          int64(el.ID),
		Name:        el.Name,
		Description: el.Description,
		Image:       el.Image,
		Points:      int64(el.Points),
		Type:        el.Type,
		Status:      el.Status,
	}
}
