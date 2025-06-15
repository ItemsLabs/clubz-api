package apiconv

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/go-openapi/strfmt"
)

// ToAppInboxSlice converts a slice of AppInbox models to the API model
func ToAppInboxSlice(inboxes schema.AppInboxSlice) []model.AppInbox {
	var result []model.AppInbox
	for _, inbox := range inboxes {
		result = append(result, ToAppInbox(inbox))
	}
	return result
}

// ToAppInbox converts an AppInbox model to the API model
func ToAppInbox(inbox *schema.AppInbox) model.AppInbox {

	var reward *model.Rewards

	var user *model.User

	// Convert Reward relationship if exists and is not nil
	if inbox.R != nil && inbox.R.Reward != nil {
		reward = ToRewards(inbox.R.Reward)
	}

	// Convert UserID relationship if exists and is not nil
	if inbox.UserID.Valid {
		user = &model.User{
			ID: inbox.UserID.String,
		}
	}

	return model.AppInbox{
		ID:          inbox.ID,
		Title:       inbox.Title,
		Description: inbox.Description,
		Status:      inbox.Status,
		Priority:    inbox.Priority,
		Category:    inbox.Category,
		ImageURL:    inbox.ImageURL.String, // Handling nullable String
		LinkURL:     inbox.LinkURL.String,  // Handling nullable String
		CreatedAt:   strfmt.DateTime(inbox.CreatedAt),
		UpdatedAt:   strfmt.DateTime(inbox.UpdatedAt),
		Read:        inbox.Read,
		User:        user, // Assign converted User
		Claimed:     inbox.Claimed,
		ClaimedAt:   strfmt.DateTime(inbox.ClamedAt.Time), // Handling nullable Time
		Clear:       inbox.Clear,
		Reward:      reward,                    // Assign converted Reward
		Match:       inbox.MatchIDID.String,    // Assign converted Match
		GameWeek:    inbox.GameWeekIDID.String, // Assign converted GameWeek
		Game:        inbox.GameID.String,       // Assign converted Game
	}
}

// ToRewards converts a schema.Reward to a model.Rewards
func ToRewards(reward *schema.Reward) *model.Rewards {
	if reward == nil {
		return nil
	}

	return &model.Rewards{
		ID:           reward.ID,
		Name:         reward.Name,
		Ball:         int64(reward.Ball),         // Casting from int or float64 to int64
		Credits:      int64(reward.Credits),      // Casting from float64 to int64
		EventTickets: int64(reward.EventTickets), // Casting from int to int64
		GameToken:    int64(reward.GameToken),    // Casting from float64 to int64
		LaptToken:    int64(reward.LaptToken),    // Casting from float64 to int64
		Shirt:        int64(reward.Shirt),        // Casting from int to int64
		SignedBall:   int64(reward.SignedBall),   // Casting from int to int64
		SignedShirt:  int64(reward.SignedShirt),  // Casting from int to int64
		KickOffPack1: int64(reward.KickoffPack1), // Casting from int to int64
		KickOffPack2: int64(reward.KickoffPack2), // Casting from int to int64
		KickOffPack3: int64(reward.KickoffPack3), // Casting from int to int64
		SeasonPack1:  int64(reward.SeasonPack1),  // Casting from int to int64
		SeasonPack2:  int64(reward.SeasonPack2),  // Casting from int to int64
		SeasonPack3:  int64(reward.SeasonPack3),  // Casting from int to int64
	}
}

// ToGameWeek converts a schema.GameWeek to a model.GameWeek
func ToGameWeek(gameWeek *schema.GameWeek) *model.GameWeek {
	if gameWeek == nil {
		return nil
	}

	return &model.GameWeek{
		ID:      gameWeek.ID,
		Name:    gameWeek.Name,
		StartAt: strfmt.DateTime(gameWeek.StartAt),
		EndAt:   strfmt.DateTime(gameWeek.EndAt),
		Status:  gameWeek.Status,
	}
}
