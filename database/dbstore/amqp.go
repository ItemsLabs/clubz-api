package dbstore

import (
	"encoding/json"

	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *DBStore) CreateAMQPEvent(exchange, typ string, data interface{}) (*schema.AmqpEvent, error) {
	if data == nil {
		data = map[string]interface{}{}
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	ev := &schema.AmqpEvent{
		Exchange: exchange,
		Type:     typ,
		Data:     string(jsonData),
	}

	return ev, ev.Insert(s.db, boil.Infer())
}
