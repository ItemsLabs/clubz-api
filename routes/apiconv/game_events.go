package apiconv

import (
	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
)

func ToGameEventSlice(events schema.GameEventSlice, actionStore database.ActionStore, dbStore database.Store) []*model.GameEvent {
	result := make([]*model.GameEvent, 0, len(events))
	for _, ev := range events {
		result = append(result, ToGameEvent(ev, actionStore, dbStore))
	}

	return result
}

func ToGameEvent(ev *schema.GameEvent, actionStore database.ActionStore, dbStore database.Store) *model.GameEvent {
	var actionName string

	action, err := actionStore.GetAction(ev.Type)
	if err != nil {
		logrus.WithError(err).WithField("id", ev.Type).Error("error on find action by id")
	} else if action == nil {
		logrus.WithField("id", ev.Type).Error("cannot find action by id")
	} else {
		actionName = action.Name
	}
	playerImageURL := ev.R.Player.ImageURL.String

	return &model.GameEvent{
		ID:              ev.ID,
		PlayerID:        ev.PlayerID,
		PlayerName:      ev.R.Player.NormalizedName.String,
		TeamID:          ev.TeamID,
		Minute:          int64(ev.Minute),
		Second:          int64(ev.Second),
		InitialScore:    ToFloatWithZero(ev.InitialScore),
		Score:           ToFloatWithZero(ev.Score),
		PoweredUp:       ev.PowerupID.Valid,
		Name:            actionName,
		CreatedAt:       strfmt.DateTime(ev.CreatedAt),
		PlayerImageURL:  playerImageURL,
		NftImageURL:     ev.NFTImage.String,
		NftMultiplier:   ev.NFTMultiplier,
		BoostMultiplier: ev.BoostMultiplier,
	}
}
