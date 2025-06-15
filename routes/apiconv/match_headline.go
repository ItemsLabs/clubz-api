package apiconv

import (
	"encoding/json"

	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
)

func ToMatchHeadlineSlice(headlines schema.MatchHeadlineSlice) []*model.MatchHeadline {
	result := make([]*model.MatchHeadline, 0, len(headlines))
	for _, h := range headlines {
		result = append(result, ToMatchHeadline(h))
	}

	return result
}

func ToMatchHeadline(headline *schema.MatchHeadline) *model.MatchHeadline {
	var images []string
	_ = json.Unmarshal([]byte(headline.Images), &images)
	if images == nil {
		images = []string{}
	}

	return &model.MatchHeadline{
		Title:       headline.Title,
		Description: headline.Description,
		Type:        headline.Type,
		ImageType:   headline.ImageType.String,
		Images:      images,
	}
}
