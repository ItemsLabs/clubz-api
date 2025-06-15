package action

import (
	"sync"

	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/database/schema"
)

type CachedActionStore struct {
	store database.Store

	actionMap map[int]*schema.Action
	mx        sync.RWMutex
}

func NewCachedActionStore(store database.Store) *CachedActionStore {
	return &CachedActionStore{
		store:     store,
		actionMap: make(map[int]*schema.Action),
	}
}

func (s *CachedActionStore) GetAction(id int) (*schema.Action, error) {
	// check within read lock
	s.mx.RLock()
	action, ok := s.actionMap[id]
	s.mx.RUnlock()

	if ok {
		return action, nil
	}

	s.mx.Lock()
	defer s.mx.Unlock()

	// check again within normal lock
	action, ok = s.actionMap[id]
	if ok {
		return action, nil
	}

	// init cache
	actions, err := s.store.GetActions()
	if err != nil {
		return nil, err
	}

	for _, act := range actions {
		s.actionMap[act.ID] = act
	}

	// look inside a cache
	action, ok = s.actionMap[id]
	if ok {
		return action, nil
	}

	return nil, nil
}
