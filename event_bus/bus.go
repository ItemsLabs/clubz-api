package event_bus

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gopkg.in/olahol/melody.v1"
	"sync"
)

type Bus interface {
	AddListener(uniqID string, m *melody.Session)
	RemoveListener(uniqID string, m *melody.Session)
	Notify(uniqID string, payload interface{}) error
}

type DefaultBus struct {
	mx        sync.Mutex
	listeners map[string]map[*melody.Session]struct{}
}

func NewDefaultBus() *DefaultBus {
	return &DefaultBus{
		listeners: map[string]map[*melody.Session]struct{}{},
	}
}

func (b *DefaultBus) AddListener(uniqID string, m *melody.Session) {
	b.mx.Lock()
	defer b.mx.Unlock()

	if _, ok := b.listeners[uniqID]; !ok {
		b.listeners[uniqID] = make(map[*melody.Session]struct{})
	}

	b.listeners[uniqID][m] = struct{}{}
}

func (b *DefaultBus) RemoveListener(uniqID string, m *melody.Session) {
	b.mx.Lock()
	defer b.mx.Unlock()

	if _, ok := b.listeners[uniqID]; ok {
		delete(b.listeners[uniqID], m)
	}
}

func (b *DefaultBus) Notify(uniqID string, payload interface{}) error {
	b.mx.Lock()
	defer b.mx.Unlock()

	// no listeners found
	sessions, ok := b.listeners[uniqID]
	if !ok {
		return nil
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for session := range sessions {
		err = session.Write(body)
		if err != nil {
			logrus.WithError(err).Error("cannot write session")
		}
	}

	return nil
}
