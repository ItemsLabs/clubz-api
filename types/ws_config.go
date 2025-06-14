package types

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/event_bus"
	"gopkg.in/olahol/melody.v1"
)

type WebsocketConfig struct {
	Manager  *melody.Melody
	EventBus event_bus.Bus
}
