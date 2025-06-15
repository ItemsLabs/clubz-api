package util

import (
	"github.com/itemslabs/clubz-api/config"
	"github.com/itemslabs/clubz-api/database"
	"github.com/palantir/stacktrace"
)

func SendPush(store database.Store, userID, matchID, title, message string, payload map[string]string) error {
	_, err := store.CreateAMQPEvent(
		config.RMQFCMExchange(),
		"push_notification",
		map[string]interface{}{
			"user_id":  userID,
			"match_id": matchID,
			"title":    title,
			"message":  message,
			"payload":  payload,
		},
	)
	if err != nil {
		return stacktrace.Propagate(err, "cannot insert amqp_event")
	}
	return nil
}
