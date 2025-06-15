package apiconv

import (
	"time"

	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
)

func ToCurrentUser(user *schema.User) *model.CurrentUser {
	banActive := false
	if user.R != nil && user.R.BanPenalties != nil {
		for _, ban := range user.R.BanPenalties {
			if !ban.EndTime.Valid || ban.EndTime.Time.After(time.Now()) {
				banActive = true
				break
			}
		}
	}
	return &model.CurrentUser{
		ID:               user.ID,
		Name:             user.Name,
		Balance:          ToFloatWithZero(user.Balance),
		PaypalEmail:      user.PaypalEmail.String,
		AvatarURL:        user.AvatarURL.String,
		Verified:         user.Verified,
		RefCode:          user.ReferralCode.String,
		RefCodeUsed:      user.ReferrerID.Valid,
		PowerupBoosts:    int64(user.BonusPowerups),
		FollowerCount:    int64(user.FollowerCount.Int),
		FollowingCount:   int64(user.FollowingCount.Int),
		GamesPlayed:      int64(user.GamesPlayed.Int),
		AvgPoints:        ToInt64PrtFromIntPtr(user.AvgPoints.Ptr()),
		AvgRank:          ToInt64PrtFromIntPtr(user.AvgRank.Ptr()),
		AvgRankPercent:   ToInt64PrtFromIntPtr(user.AvgRankPercent.Ptr()),
		Moderator:        user.Moderator,
		Banned:           banActive,
		Premium:          user.Premium,
		SubscriptionTier: ConvertSubscriptionTier(user.SubscriptionTier),
		Influencer:       user.Influencer,
		Email:            user.Email.String,
	}
}

func ToUser(user *schema.User, isFollowing bool) *model.User {
	return &model.User{
		ID:               user.ID,
		Name:             user.Name,
		AvatarURL:        user.AvatarURL.String,
		Verified:         user.Verified,
		FollowerCount:    int64(user.FollowerCount.Int),
		FollowingCount:   int64(user.FollowingCount.Int),
		GamesPlayed:      int64(user.GamesPlayed.Int),
		AvgPoints:        ToInt64PrtFromIntPtr(user.AvgPoints.Ptr()),
		AvgRank:          ToInt64PrtFromIntPtr(user.AvgRank.Ptr()),
		AvgRankPercent:   ToInt64PrtFromIntPtr(user.AvgRankPercent.Ptr()),
		IsFollowing:      isFollowing,
		Premium:          user.Premium,
		SubscriptionTier: ConvertSubscriptionTier(user.SubscriptionTier),
		Influencer:       user.Influencer,
		Email:            user.Email.String,
	}
}

func ToShortUser(user *schema.User) *model.ShortUser {
	return &model.ShortUser{
		ID:               user.ID,
		Name:             user.Name,
		AvatarURL:        user.AvatarURL.String,
		Verified:         user.Verified,
		Premium:          user.Premium,
		SubscriptionTier: ConvertSubscriptionTier(user.SubscriptionTier),
		Influencer:       user.Influencer,
		Moderator:        user.Moderator,
		Email:            user.Email.String,
	}
}

func ToShortUserArray(users schema.UserSlice) model.ShortUserArray {
	var result = make(model.ShortUserArray, 0, len(users))
	for _, user := range users {
		result = append(result, ToShortUser(user))
	}

	return result
}

func ConvertSubscriptionTier(val int) model.SubscriptionTier {
	switch val {
	case database.SubscriptionTierPremium:
		return model.SubscriptionTierPremium
	case database.SubscriptionTierLite:
		return model.SubscriptionTierLite
	default:
		return model.SubscriptionTierNone
	}
}
