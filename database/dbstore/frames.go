package dbstore

import (
	"database/sql"

	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetFrames retrieves all frames, ordered by points in descending order.
func (s *DBStore) GetFrames() (schema.FrameSlice, error) {
	return schema.Frames(
		qm.OrderBy("points desc"),
	).All(s.db)
}

// GetFrameNameByID returns the name of a frame given its ID.
// It returns an empty string and an error if no frame is found for the given ID.
func (s *DBStore) GetFrameNameByID(frameID int) (string, error) {
	frame, err := schema.FindFrame(s.db, frameID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return frame.Name, nil
}

// GetFrameByID retrieves a frame by its ID.
func (s *DBStore) GetFrameByID(frameID string) (*schema.Frame, error) {
	return schema.Frames(
		qm.Where("id = ?", frameID),
	).One(s.db)
}

// DeleteFrame deletes a frame given its ID.
// It also deletes associated records from the user_frames table if necessary.
func (s *DBStore) DeleteFrame(frameID string) error {
	_, err := schema.UserFrames(
		qm.Where("frame_id = ?", frameID),
	).DeleteAll(s.db)
	if err != nil {
		return err
	}

	frame, err := schema.Frames(
		qm.Where("id = ?", frameID),
	).One(s.db)
	if err != nil {
		return err
	}

	_, err = frame.Delete(s.db)
	return err
}

// UpdateFrame updates a frame's fields given its ID and an options struct.
func (s *DBStore) UpdateFrame(frameID string, options *model.UpdateFrameRequest) (*schema.Frame, error) {
	frame, err := schema.Frames(qm.Where("id = ?", frameID)).One(s.db)
	if err != nil {
		return nil, err
	}

	frame.Name = options.Name
	frame.Description = options.Description
	frame.Image = options.Image
	frame.Points = int(options.Points)
	frame.Type = options.Type
	frame.Status = options.Status

	_, err = frame.Update(s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return frame, nil
}

// CreateFrames inserts multiple frames into the database and returns the newly inserted frames.
func (s *DBStore) CreateFrames(frames *model.CreateFramesRequest) ([]schema.Frame, error) {
	var insertedFrames []schema.Frame
	var maxID int
	err := s.db.QueryRow("SELECT MAX(id) FROM frames").Scan(&maxID)
	if err != nil {
		return nil, err
	}

	for _, frame := range *frames {
		maxID++

		schemaFrame := schema.Frame{
			ID:          maxID,
			Name:        *frame.Name,
			Description: *frame.Description,
			Image:       *frame.Image,
			Points:      int(*frame.Points),
			Type:        *frame.Type,
			Status:      *frame.Status,
		}

		err := schemaFrame.Insert(s.db, boil.Infer())
		if err != nil {
			return nil, err
		}

		insertedFrames = append(insertedFrames, schemaFrame)
	}

	return insertedFrames, nil
}

func (s *DBStore) GetDefaultFrames() (schema.FrameSlice, error) {
	return schema.Frames(
		qm.Where("type = 'default'"),
	).All(s.db)
}
