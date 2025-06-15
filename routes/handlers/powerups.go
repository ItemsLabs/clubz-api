package handlers

import (
	"github.com/itemslabs/clubz-api/config"
	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/apiconv"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

type PowerUpService struct {
}

// PowerUpList godoc
// @Summary List all powerups
// @Description List all powerups
// @ID powerups-list
// @Produce json
// @Success 200 {object} model.PowerUp
// @Router /:id/powerups [get]
func (e *Env) PowerUpList(c echo.Context) error {
	puList, err := e.Store.GetPowerUps()
	if err != nil {
		return err
	}

	actions, err := e.Store.GetPowerUpActions()
	if err != nil {
		return err
	}

	// distribute actions by powerups
	var puActions = map[int]schema.ActionSlice{}
	for _, act := range actions {
		if _, ok := puActions[act.ActionID]; !ok {
			puActions[act.ActionID] = schema.ActionSlice{}
		}

		puActions[act.PowerupID] = append(puActions[act.PowerupID], act.R.Action)
	}

	return e.RespondSuccess(c, apiconv.ToPowerUpSlice(puList, puActions))
}

// ApplyPowerUp godoc
// @Summary Apply PowerUp
// @Description Apply PowerUp
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Game ID"
// @Param body body model.ApplyPowerUpRequest true "Apply PowerUp request"
// @Success 204 {object} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /games/{id}/powerup [post]
func (e *Env) ApplyPowerUp(c echo.Context) error {
	var in model.ApplyPowerUpRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}

	// check incoming data
	// double-up id
	puID := int(*in.PowerupID)
	var pos = int(in.LineupPosition)

	// select pu
	powerup, err := e.Store.GetPowerUpByID(puID)
	if err != nil {
		return err
	}

	var powerupType int
	if powerup.Type.Valid && powerup.Type.Int > 0 {
		powerupType = powerup.Type.Int
	}

	if (pos < 1 || pos > config.MaxPicks()) && (powerupType != database.PowerUpTypeSpecial) {
		return ErrInvalidPowerUpPosition
	}

	// check game belong to user
	gameID := c.Param("id")
	ok, err := e.Store.IsGameBelongToUser(gameID, userID(c))
	if err != nil {
		return err
	}

	if !ok {
		return ErrObjectNotFound
	}

	return e.Store.Transaction(func(s database.Store) error {
		game, err := s.LockGame(gameID)
		if err != nil {
			return err
		}

		user, err := s.LockUser(userID(c))
		if err != nil {
			return err
		}

		if user.Balance-powerup.Cost < 0 {
			return ErrNotEnoughBalanceForPowerup
		}

		match, err := s.GetMatchByID(game.MatchID)
		if err != nil {
			return err
		}

		if match.Status != database.MatchStatusGame {
			return ErrInvalidMatchStatusForPowerUp
		}

		powerups, err := s.GetGamePowerUps(gameID)
		if err != nil {
			return err
		}

		var bonus = false
		if pos > 0 {
			// check for active power ups for this position
			var countForPosition = 0
			for _, pu := range powerups {
				if pu.Position == pos {
					if !pu.EndedAt.Valid {
						return ErrAlreadyActivePowerUpForPosition
					}
					countForPosition++
				}
			}

			// we already have power-up for user
			// if countForPosition == 1 {
			// 	if game.Premium || game.SubscriptionTier == database.SubscriptionTierPremium ||
			// 		game.SubscriptionTier == database.SubscriptionTierLite {

			// 		// subscription users allow to use bonus power-up in every game
			// 		bonus = true
			// 	} else if user.BonusPowerups > 0 {
			// 		user.BonusPowerups--
			// 		if _, err = s.UpdateBonusPowerUps(user); err != nil {
			// 			return err
			// 		}
			// 		bonus = true
			// 	} else {
			// 		return ErrNoPowerUpLeft
			// 	}
			// } else if countForPosition > 1 {
			// 	return ErrNoPowerUpLeft
			// }
		}

		// calculate duration and multiplier
		switch in.Option {
		case model.PowerupOptionX35:
			powerup.Duration = 5 * 60
			powerup.Multiplier = 3
		case model.PowerupOptionX1515:
			powerup.Duration = 15 * 60
			powerup.Multiplier = 1.5
		default:
			// powerup.Duration = 10 * 60
			// powerup.Multiplier = 2 // Should use the set multiplier at DB level.
		}

		_, err = s.CreateGamePowerUp(&schema.GamePowerup{
			ID:         uuid.New().String(),
			GameID:     gameID,
			PowerupID:  puID,
			Position:   pos,
			Duration:   powerup.Duration,
			Multiplier: powerup.Multiplier,
			Minute:     match.Minute,
			Second:     match.Second,
			Bonus:      bonus,
		})

		if err != nil {
			return err
		}

		// Create the transaction to update the Balance properly
		if _, err := s.CreateTransaction(&schema.Transaction{
			ID:         uuid.NewString(),
			UserID:     user.ID,
			Amount:     -powerup.Cost,
			ObjectType: database.TransactionTypeVirtual,
			MatchID:    null.StringFrom(game.MatchID),
			Quantity:   int64(1),
		}); err != nil {
			return err
		}

		// notify that was updated
		_, err = s.CreateAMQPEvent(config.AmqpGamesExchange(), "game_updated", map[string]interface{}{
			"game_id":  gameID,
			"match_id": match.ID,
		})
		if err != nil {
			return err
		}

		return e.RespondNoContent(c)
	})
}

// GetGamePowerUps godoc
// @Summary Get game powerups
// @Description Get game powerups
// @ID get-game-powerups
// @Produce json
// @Param id path string true "Game ID"
// @Success 200 {array} model.GamePowerup
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /games/{id}/powerups [get]
func (e *Env) GetGamePowerUps(c echo.Context) error {
	gameID := c.Param("id")
	powerups, err := e.Store.GetGamePowerUps(gameID)
	if err != nil {
		return err
	}
	return e.RespondSuccess(c, powerups)
}
