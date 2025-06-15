package util

import (
	"time"

	"github.com/itemslabs/clubz-api/database/schema"
)

func UserHasBanPenalties(user *schema.User) bool {
	if user.R != nil && user.R.BanPenalties != nil {
		for _, ban := range user.R.BanPenalties {
			if !ban.EndTime.Valid || ban.EndTime.Time.After(time.Now()) {
				return true
			}
		}
	}

	return false
}
