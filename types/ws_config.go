package types

import (
	"github.com/itemslabs/clubz-api/event_bus"
	"gopkg.in/olahol/melody.v1"
)

type WebsocketConfig struct {
	Manager  *melody.Melody
	EventBus event_bus.Bus
}
