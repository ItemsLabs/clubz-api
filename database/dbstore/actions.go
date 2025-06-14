package dbstore

import (
	"database/sql"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) GetActions() (schema.ActionSlice, error) {
	return schema.Actions(
		qm.OrderBy("score desc"),
	).All(s.db)
}

// GetActionNameByActionID returns the name of an action given its ID.
// It returns an empty string and an error if no action is found for the given ID.
func (s *DBStore) GetActionNameByActionID(actionID int) (string, error) {
	// Ensure s.db is properly initialized and connected to your database
	action, err := schema.FindAction(s.db, actionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // or return an error indicating that no action was found
		}
		return "", err
	}

	return action.Name, nil
}

func (s *DBStore) GetActionByID(actionID string) (*schema.Action, error) {
	return schema.Actions(
		qm.Where("id = ?", actionID),
	).One(s.db)
}

func (s *DBStore) DeleteAction(actionID string) error {
	_, err := schema.PowerupActions(
		qm.Where("action_id = ?", actionID),
	).DeleteAll(s.db)
	if err != nil {
		return err
	}

	action, err := schema.Actions(
		qm.Where("id = ?", actionID),
	).One(s.db)

	if err != nil {
		return err
	}

	_, err = action.Delete(s.db)
	return err
}

func (s *DBStore) UpdateAction(actionID string, options *model.UpdateActionRequest) (*schema.Action, error) {
	action, err := schema.Actions(qm.Where("id = ?", actionID)).One(s.db)
	if err != nil {
		return nil, err
	}

	action.Name = options.Name
	action.Description = options.Description
	action.Score = options.Score

	_, err = action.Update(s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return action, nil
}

func (s *DBStore) CreateActions(actions *model.CreateActionsRequest) ([]schema.Action, error) {
	var insertedActions []schema.Action
	var maxID int
	err := s.db.QueryRow("SELECT MAX(id) FROM actions").Scan(&maxID)
	if err != nil {
		return nil, err
	}

	for _, action := range *actions {
		maxID++

		schemaAction := schema.Action{
			ID:          maxID,
			Name:        action.Name,
			Description: action.Description,
			Score:       action.Score,
		}

		err := schemaAction.Insert(s.db, boil.Infer())
		if err != nil {
			return nil, err
		}

		insertedActions = append(insertedActions, schemaAction)
	}

	return insertedActions, nil

}
