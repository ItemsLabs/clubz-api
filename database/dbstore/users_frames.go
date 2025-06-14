package dbstore

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetUserFrames retrieves all frames associated with a user.
func (s *DBStore) GetUserFrames(userID string) (schema.UserFrameSlice, error) {
	return schema.UserFrames(
		qm.Where("user_id = ?", userID),
	).All(s.db)
}

// GetUserFrameByID retrieves a user frame by its ID.
func (s *DBStore) GetUserFrameByID(userFrameID int) (*schema.UserFrame, error) {
	return schema.UserFrames(
		qm.Where("id = ?", userFrameID),
	).One(s.db)
}

// DeleteUserFrame deletes a user frame by its ID.
func (s *DBStore) DeleteUserFrame(userFrameID int) error {
	userFrame, err := schema.UserFrames(
		qm.Where("id = ?", userFrameID),
	).One(s.db)
	if err != nil {
		return err
	}
	_, err = userFrame.Delete(s.db)
	return err
}

// CreateUserFrame creates a user frame for a specific user.
func (s *DBStore) CreateUserFrame(userID string, frameID int) (*schema.UserFrame, error) {
	userFrame := &schema.UserFrame{
		UserID:  userID,
		FrameID: frameID,
	}

	err := userFrame.Insert(s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return userFrame, nil
}

func (s *DBStore) ChangeAllUsersFramesToUnselected() error {
	_, err := s.db.Exec("UPDATE user_frames SET selected = false")
	return err
}

func (s *DBStore) ChangeUserFrameToSelected(userFrameID int) error {
	_, err := s.db.Exec("UPDATE user_frames SET selected = true WHERE id = ?", userFrameID)
	return err
}
