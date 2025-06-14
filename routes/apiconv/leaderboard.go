package apiconv

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
)

func ToLeaderBoardEntry(el *schema.MatchLeaderboard) *model.LeaderboardEntry {
	leaderboard := &model.LeaderboardEntry{
		UserID:           el.UserID,
		UserName:         el.R.User.Name,
		UserAvatarURL:    el.R.User.AvatarURL.String,
		Score:            ToFloatWithZero(el.Score.Float64),
		Premium:          el.R.User.Premium,
		SubscriptionTier: ConvertSubscriptionTier(el.R.User.SubscriptionTier),
		Position:         int64(el.Position.Int),
	}
	if el.R.Transaction != nil {
		leaderboard.Reward = ToFloatWithZero(el.R.Transaction.Amount)
	}
	return leaderboard
}
