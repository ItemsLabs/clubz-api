package listeners

import (
	"context"
	"encoding/json"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/amqp"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/event_bus"
	"github.com/sirupsen/logrus"
)

type gameListener struct {
	bus event_bus.Bus
}

func newGameListener(bus event_bus.Bus) *gameListener {
	return &gameListener{bus: bus}
}

func (l *gameListener) OnConnect() {
	logrus.Info("Connected to AMQP server.")
}

func (l *gameListener) OnReconnect() {
	logrus.Info("Reconnected to AMQP server.")
}

func (l *gameListener) OnDeliveryReceived(sess amqp.Session, d amqp.Delivery) {
	switch d.Type {
	case "update":
		l.handleUpdateEvent(d)
	case "leaderboard_updated":
		l.handleLeaderboardUpdatedEvent(d)
	default:
		logrus.WithField("type", d.Type).Error("Unhandled message type received")
		_ = d.Nack(false, false)
	}
}

func (l *gameListener) handleUpdateEvent(d amqp.Delivery) {
	var payload struct {
		UserID           string `json:"user_id"`
		GameID           string `json:"game_id"`
		PlayerID         string `json:"player_id"`
		GameEvent        string `json:"game_event"`
		GamePowerUp      string `json:"game_powerup"`
		GameInitialScore string `json:"game_initial_score"`
		GameScore        string `json:"game_score"`
		PlayerName       string `json:"player_name"`
		TeamID           string `json:"team_id"`
		PlayerImage      string `json:"player_image"`
		GameEventId      string `json:"game_event_id"`
		EventMinute      string `json:"event_minute"`
		Eventseconds     string `json:"event_second"`
		MatchID          string `json:"match_id"`
		NormalizedName   string `json:"normalized_name"`
		NFTMultiplier    string `json:"nft_multiplier"`
		BoostMultiplier  string `json:"boost_multiplier"`
		NFTImage         string `json:"nft_image"`
	}

	if err := json.Unmarshal(d.Body, &payload); err != nil {
		logrus.WithError(err).Error("Failed to unmarshal update event")
		_ = d.Nack(false, false)
		return
	}

	msg := map[string]interface{}{
		"type":               "new",
		"game_id":            payload.GameID,
		"user_id":            payload.UserID,
		"player_id":          payload.PlayerID,
		"game_event":         payload.GameEvent,
		"game_powerup":       payload.GamePowerUp,
		"game_initial_score": payload.GameInitialScore,
		"game_score":         payload.GameScore,
		"player_name":        payload.PlayerName,
		"team_id":            payload.TeamID,
		"player_image":       payload.PlayerImage,
		"game_event_id":      payload.GameEventId,
		"event_minute":       payload.EventMinute,
		"event_second":       payload.Eventseconds,
		"match_id":           payload.MatchID,
		"normalized_name":    payload.NormalizedName,
		"nft_multiplier":     payload.NFTMultiplier,
		"boost_multiplier":   payload.BoostMultiplier,
		"nft_image":          payload.NFTImage,
	}
	body, _ := json.Marshal(msg)
	_ = l.bus.Notify(payload.UserID, body)
	_ = d.Ack(false)
}

func (l *gameListener) handleLeaderboardUpdatedEvent(d amqp.Delivery) {
	type leaderboardEvent struct {
		MatchID string `json:"match_id"`
		Entries []struct {
			ID               int      `json:"id"`
			Position         *int     `json:"position"`
			UserName         string   `json:"user_name"`
			UserID           string   `json:"user_id"`
			UserAvatarURL    *string  `json:"user_avatar_url"`
			Score            *float64 `json:"score"`
			Premium          bool     `json:"premium"`
			SubscriptionTier string   `json:"subscription_tier"`
			Influencer       bool     `json:"influencer"`
		}
	}

	var payload leaderboardEvent
	if err := json.Unmarshal(d.Body, &payload); err != nil {
		logrus.WithError(err).Error("Failed to unmarshal leaderboard updated event")
		_ = d.Nack(false, false)
		return
	}

	msg := map[string]interface{}{
		"type":      "leaderboard_updated",
		"match_id":  payload.MatchID,
		"entries":   payload.Entries,
		"timestamp": d.Timestamp.String(),
	}
	body, _ := json.Marshal(msg)
	_ = l.bus.Notify(payload.MatchID, body)
	_ = d.Ack(false)
}

func StartGameEventsListener(ctx context.Context, url, exchange, queue string, bus event_bus.Bus) {
	amqp.NewSubscriber(
		ctx,
		url,
		newGameListener(bus),
		amqp.ExclusiveQueueDefiner(queue, exchange),
		amqp.PrefetchConsumerDefiner(queue, 100),
	)
}
