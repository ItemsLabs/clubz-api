package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/config"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/apiconv"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/palantir/stacktrace"
	"github.com/volatiletech/null/v8"
)

type PlayerWithNFT struct {
	Player schema.Player
	NFT    string
}

// GamesList godoc
// @Summary List all games
// @Description List all games
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param debug48HoursAgo query string false "Debug 48 hours ago"
// @Success 200 {object} []model.Game
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Router /games [get]
func (e *Env) GamesList(c echo.Context) error {
	now := time.Now()
	lowDelta := now.Add(-config.MatchDisplayLowDelta())

	if c.QueryParam("debug48HoursAgo") != "" {
		lowDelta = now.Add(-48 * time.Hour)
	}

	games, err := e.Store.GetGamesLaterThan(userID(c), lowDelta)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToGameSlice(games))
}

// GetGameByID godoc
// @Summary Get game by ID
// @Description Get game by ID
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Game ID"
// @Success 200 {object} model.Game
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /games/{id} [get]
func (e *Env) GetGameByID(c echo.Context) error {
	game, err := e.Store.GetGameByID(c.Param("id"), userID(c))
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToGame(game))
}

// GetUserGame godoc
// @Summary Get user game
// @Description Get user game
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Game ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} model.Game
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /games/{id}/user/{user_id} [get]
func (e *Env) GetUserGame(c echo.Context) error {
	currentUser, err := e.Store.GetUserByID(userID(c))
	if err != nil {
		return err
	}

	if currentUser.SubscriptionTier != database.SubscriptionTierPremium {
		return ErrNotAuthorized
	}

	matchID, targetUserID := c.Param("id"), c.Param("user_id")

	game, err := e.Store.GetGameByUserIDMatchID(targetUserID, matchID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToGame(game))
}

// JoinGame godoc
// @Summary Join game
// @Description Join game
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Game ID"
// @Success 200 {object} model.Game
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /games/{id}/join [post]
func (e *Env) JoinGame(c echo.Context) error {
	log.Println("JoinGame called")
	var in model.GameJoinRequest

	// Parse request body
	if err := e.ParseBody(c, &in); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	log.Println("Request body parsed successfully")

	userID := userID(c)
	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		log.Printf("Error retrieving user by ID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user")
	}
	log.Printf("User retrieved: %v", user)

	isBanned := false
	if user.R != nil && user.R.BanPenalties != nil {
		for _, ban := range user.R.BanPenalties {
			if !ban.EndTime.Valid || ban.EndTime.Time.After(time.Now()) {
				isBanned = true
			}
		}
	}
	if isBanned {
		log.Println("User is banned")
		return echo.NewHTTPError(http.StatusForbidden, "User is banned")
	}

	if len(in.Picks) != 4 {
		log.Println("Invalid number of picks")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid number of picks")
	}
	log.Println("Valid number of picks")

	assignedCardPacks, err := e.Store.GetAssignedCardPacksByUserID(userID)
	if err != nil {
		log.Printf("Error retrieving assigned card packs by user ID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve assigned card packs")
	}
	log.Printf("Assigned card packs retrieved: %v", assignedCardPacks)

	var mintedPlayers []map[string]interface{}
	// Use the subgraph to get the balance and NFT IDs
	nftIDs, err := getNFTDataFromCurl(user.WalletAddress.String)
	if err != nil {
		log.Printf("Error fetching NFT data from subgraph for address %s: %v", user.WalletAddress.String, err)

	}

	playersWithNFTs, err := e.Store.GetAssignedPlayersByNFTIDs(nftIDs)
	if err != nil {
		log.Printf("Error retrieving assigned players by NFT IDs: %v", err)

	}

	for _, playerWithNFT := range playersWithNFTs {
		playerNFTID := playerWithNFT.Player.PlayerNFTID.String
		nftDetails, err := e.Store.GetNFTBucketByID(playerNFTID)
		if err != nil {
			log.Printf("Error retrieving NFT bucket by ID: %v", err)
			continue
		}

		mintedPlayers = append(mintedPlayers, map[string]interface{}{
			"player": playerWithNFT.Player,
			"nft":    playerWithNFT.NFT,
			"name":   nftDetails.Name,
		})
	}
	for _, cardPack := range assignedCardPacks {
		log.Printf("Processing card pack ID %s for user ID %s", cardPack.ID, userID)
		if cardPack.OpenedAt.Valid {
			if !cardPack.Revealed {
				log.Printf("Processing card pack opened within 48 hours: %v", cardPack.ID)
				cardPackWithNFTDetails, err := e.Store.GetAssignedCardPackByID(cardPack.ID)
				if err != nil {
					log.Printf("Error retrieving assigned card pack by ID: %v", err)
					continue
				}

				cardIDs := cardPackWithNFTDetails.CardIds
				if cardIDs.IsZero() {
					log.Println("Card IDs are zero")
					continue
				}

				var cardUUIDs []string
				if err := cardIDs.Unmarshal(&cardUUIDs); err != nil {
					log.Printf("Error unmarshaling card IDs: %v", err)
					continue
				}

				playersWithNFTs, err := e.Store.GetAssignedPlayersWithNFTDetails(cardUUIDs)
				if err != nil {
					log.Printf("Error retrieving assigned players with NFT details: %v", err)
					continue
				}

				for _, playerWithNFT := range playersWithNFTs {
					playerNFTID := playerWithNFT.Player.PlayerNFTID.String
					nftDetails, err := e.Store.GetNFTBucketByID(playerNFTID)
					if err != nil {
						log.Printf("Error retrieving NFT bucket by ID: %v", err)
						continue
					}

					mintedPlayers = append(mintedPlayers, map[string]interface{}{
						"player": playerWithNFT.Player,
						"nft":    playerWithNFT.NFT,
						"name":   nftDetails.Name,
					})
				}
			} else {
				log.Printf("CardPack already minted: %v", cardPack.ID)

			}
		}
	}
	log.Printf("Minted players: %v", mintedPlayers)

	// Select match and players
	match, err := e.Store.GetMatchByID(*in.MatchID)
	if err != nil {
		log.Printf("Error retrieving match by ID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve match")
	}
	log.Printf("Match retrieved: %v", match)

	var players schema.PlayerSlice
	for _, p := range in.Picks {
		player, err := e.Store.GetPlayerByID(p)
		if err != nil {
			log.Printf("Error retrieving player by ID: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve player")
		}
		players = append(players, player)
	}
	log.Printf("Players retrieved: %v", players)

	if match.MatchType == database.MatchTypeUnknown {
		log.Println("Invalid match type")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid match type")
	}

	// Check whether game already exists
	activeGameID, err := e.Store.GetActiveGameIDForMatch(match.ID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Game not exists, create new one
			log.Println("Creating new game")
			err = e.Store.Transaction(func(s database.Store) error {
				// Find proper game status
				var gameStatus = database.GameStatusWaiting
				if match.Status == database.MatchStatusGame {
					gameStatus = database.GameStatusGameplay
				}

				// Calculate game num
				gameNum, err := s.GetNumberOfGames(userID)
				if err != nil {
					log.Printf("Error getting number of games: %v", err)
					return err
				}

				// Create game
				game, err := s.CreateGame(&schema.Game{
					ID:               uuid.New().String(),
					UserID:           userID,
					MatchID:          match.ID,
					Status:           gameStatus,
					Premium:          user.Premium,
					SubscriptionTier: user.SubscriptionTier,
					Num:              null.IntFrom(gameNum + 1),
					SportID:          match.SportID,
					Notified:         false,
				})
				if err != nil {
					log.Printf("Error creating game: %v", err)
					return err
				}

				// Create picks
				for idx, player := range players {
					var assignedPlayerID null.String
					// Filter and find the rarest player
					var matchingPlayers []map[string]interface{}
					for _, mintedPlayer := range mintedPlayers {
						if mintedPlayer["nft"].(*schema.NFTBucket).OptaID == player.ImportID {
							matchingPlayers = append(matchingPlayers, mintedPlayer)
						}
					}

					if len(matchingPlayers) > 0 {
						rarestPlayer := findRarestPlayer(matchingPlayers)
						assignedPlayerID = null.StringFrom(rarestPlayer["player"].(*schema.AssignedPlayer).ID)
					}

					gamePick := &schema.GamePick{
						ID:               uuid.New().String(),
						GameID:           game.ID,
						PlayerID:         player.ID,
						Position:         idx + 1,
						Minute:           match.Minute,
						Second:           match.Second,
						UserSwapped:      true,
						AssignedPlayerID: assignedPlayerID,
					}

					_, err = s.CreateGamePick(gamePick)
					if err != nil {
						log.Printf("Error creating game pick for player %s: %v", player.ID, err)
						return err
					}
				}

				// Notify that game was created
				_, err = s.CreateAMQPEvent(config.AmqpGamesExchange(), "game_updated", map[string]interface{}{
					"game_id":  game.ID,
					"match_id": match.ID,
				})
				if err != nil {
					log.Printf("Error creating AMQP event: %v", err)
					return err
				}

				activeGameID = game.ID

				return nil
			})
			if err != nil {
				log.Printf("Error creating game: %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create game")
			}
		} else {
			log.Printf("Error retrieving active game: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve active game")
		}
	}

	// Game exists, query from db
	game, err := e.Store.GetGameByID(activeGameID, userID)
	if err != nil {
		log.Printf("Error retrieving game by ID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve game")
	}
	log.Printf("Game retrieved: %v", game)

	return e.RespondSuccess(c, apiconv.ToGame(game))
}

// SwapPlayer godoc
// @Summary Swap player
// @Description Swap player
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Game ID"
// @Param body body model.GameSwapRequest true "Game swap request"
// @Success 200 {object} model.GamePick
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /games/{id}/swap [post]
func (e *Env) SwapPlayer(c echo.Context) error {

	return e.Substitution(c)

	// TODO: remove this implementation from the handler and show this if still someone is calling this
	// 	return &APIError{
	// 		Code:       108,
	// 		Message:    "This endpoint is deprecated. Use /games/{id}/substitution instead.",
	// 		StatusCode: http.StatusSeeOther,
	// 	}

	// var in model.GameSwapRequest

	// if err := e.ParseBody(c, &in); err != nil {
	// 	return err
	// }

	// // find game, pick and player
	// gameID := c.Param("id")
	// game, err := e.Store.GetGameByID(gameID, userID(c))
	// if err != nil {
	// 	return stacktrace.Propagate(err,
	// 		fmt.Sprintf("GetGameByID failed (game_id: %s, user_id: %s)", gameID, userID(c)))
	// }

	// pick, err := e.Store.GetGamePickByID(*in.PlayerPickID, game.ID)
	// if err != nil {
	// 	return stacktrace.Propagate(err,
	// 		fmt.Sprintf("GetGamePickByID failed (pick_id: %s, game_id: %s)", *in.PlayerPickID, game.ID))
	// }

	// player, err := e.Store.GetPlayerByID(*in.NewPlayerID)
	// if err != nil {
	// 	return stacktrace.Propagate(err,
	// 		fmt.Sprintf("GetPlayerByID faiked (id: %s)", *in.NewPlayerID))
	// }

	// if pick.EndedAt.Valid {
	// 	return ErrPickAlreadyEnded
	// }

	// // check whether we already have active pick for this player
	// activePicks, err := e.Store.GetActiveGamePicksWithPlayer(game.ID, player.ID)
	// if err != nil {
	// 	return stacktrace.Propagate(err,
	// 		fmt.Sprintf("GetActiveGamePicksWithPlayer failed (game_id: %s, player_id: %s)", game.ID, player.ID))
	// }

	// if len(activePicks) > 0 {
	// 	return ErrYouAlreadyPickedThisPlayer
	// }

	// now := time.Now()

	// var resultPick *schema.GamePick
	// if err = e.Store.Transaction(func(s database.Store) error {
	// 	// finalize current pick
	// 	pick.EndedAt = null.TimeFrom(now)
	// 	if err := s.UpdateGamePickEndedAt(pick); err != nil {
	// 		return err
	// 	}

	// 	// create new pick
	// 	newPick, err := s.CreateGamePick(&schema.GamePick{
	// 		ID:          uuid.New().String(),
	// 		CreatedAt:   now,
	// 		GameID:      game.ID,
	// 		PlayerID:    player.ID,
	// 		Position:    pick.Position,
	// 		Minute:      game.R.Match.Minute,
	// 		Second:      game.R.Match.Second,
	// 		UserSwapped: game.Status == database.GameStatusGameplay,
	// 	})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	resultPick = newPick

	// 	return nil
	// }); err != nil {
	// 	return err
	// }

	// return e.RespondSuccess(c, apiconv.ToGamePick(resultPick,
	// 	util.IsPlayerPlaying(player.ID, game.R.Match.R.MatchPlayers)))
}

// Substitution godoc
// @Summary Substitutes a player
// @Description Substitution of a player
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Game ID"
// @Param body body model.GameSwapRequest true "Game substitution (Swap)request"
// @Success 200 {object} model.GamePick
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /games/{id}/substitution [post]
func (e *Env) Substitution(c echo.Context) error {
	var in model.GameSwapRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}
	userID := userID(c)
	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		log.Printf("Error retrieving user by ID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user")
	}
	log.Printf("User retrieved: %v", user)
	// find game, pick and player
	gameID := c.Param("id")
	game, err := e.Store.GetGameByID(gameID, userID)

	if game.Status != database.GameStatusWaiting && game.Status != database.GameStatusGameplay {
		return ErrSubstitutionMatchInvalidState
	}

	if err != nil {
		return stacktrace.Propagate(err,
			fmt.Sprintf("GetGameByID failed (game_id: %s, user_id: %s)", gameID, userID))
	}

	pick, err := e.Store.GetGamePickByID(*in.PlayerPickID, game.ID)
	if err != nil {
		return stacktrace.Propagate(err,
			fmt.Sprintf("GetGamePickByID failed (pick_id: %s, game_id: %s)", *in.PlayerPickID, game.ID))
	}

	player, err := e.Store.GetPlayerByID(*in.NewPlayerID)
	if err != nil {
		return stacktrace.Propagate(err,
			fmt.Sprintf("GetPlayerByID failed (id: %s)", *in.NewPlayerID))
	}

	if pick.EndedAt.Valid {
		return ErrPickAlreadyEnded
	}

	// check whether we already have active pick for this player
	activePicks, err := e.Store.GetActiveGamePicksWithPlayer(game.ID, player.ID)
	if err != nil {
		return stacktrace.Propagate(err,
			fmt.Sprintf("GetActiveGamePicksWithPlayer failed (game_id: %s, player_id: %s)", game.ID, player.ID))
	}

	if len(activePicks) > 0 {
		return ErrYouAlreadyPickedThisPlayer
	}

	now := time.Now()
	assignedCardPacks, err := e.Store.GetAssignedCardPacksByUserID(userID)
	if err != nil {
		log.Printf("Error retrieving assigned card packs by user ID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve assigned card packs")
	}
	log.Printf("Assigned card packs retrieved: %v", assignedCardPacks)

	var mintedPlayers []map[string]interface{}
	// Use the subgraph to get the balance and NFT IDs
	nftIDs, err := getNFTDataFromCurl(user.WalletAddress.String)
	if err != nil {
		log.Printf("Error fetching NFT data from subgraph for address %s: %v", user.WalletAddress.String, err)

	}

	playersWithNFTs, err := e.Store.GetAssignedPlayersByNFTIDs(nftIDs)
	if err != nil {
		log.Printf("Error retrieving assigned players by NFT IDs: %v", err)

	}

	for _, playerWithNFT := range playersWithNFTs {
		playerNFTID := playerWithNFT.Player.PlayerNFTID.String
		nftDetails, err := e.Store.GetNFTBucketByID(playerNFTID)
		if err != nil {
			log.Printf("Error retrieving NFT bucket by ID: %v", err)
			continue
		}

		mintedPlayers = append(mintedPlayers, map[string]interface{}{
			"player": playerWithNFT.Player,
			"nft":    playerWithNFT.NFT,
			"name":   nftDetails.Name,
		})
	}
	for _, cardPack := range assignedCardPacks {
		log.Printf("Processing card pack ID %s for user ID %s", cardPack.ID, userID)
		if cardPack.OpenedAt.Valid {
			if !cardPack.Revealed {
				log.Printf("Processing card pack opened within 48 hours: %v", cardPack.ID)
				cardPackWithNFTDetails, err := e.Store.GetAssignedCardPackByID(cardPack.ID)
				if err != nil {
					log.Printf("Error retrieving assigned card pack by ID: %v", err)
					continue
				}

				cardIDs := cardPackWithNFTDetails.CardIds
				if cardIDs.IsZero() {
					log.Println("Card IDs are zero")
					continue
				}

				var cardUUIDs []string
				if err := cardIDs.Unmarshal(&cardUUIDs); err != nil {
					log.Printf("Error unmarshaling card IDs: %v", err)
					continue
				}

				playersWithNFTs, err := e.Store.GetAssignedPlayersWithNFTDetails(cardUUIDs)
				if err != nil {
					log.Printf("Error retrieving assigned players with NFT details: %v", err)
					continue
				}

				for _, playerWithNFT := range playersWithNFTs {
					playerNFTID := playerWithNFT.Player.PlayerNFTID.String
					nftDetails, err := e.Store.GetNFTBucketByID(playerNFTID)
					if err != nil {
						log.Printf("Error retrieving NFT bucket by ID: %v", err)
						continue
					}

					mintedPlayers = append(mintedPlayers, map[string]interface{}{
						"player": playerWithNFT.Player,
						"nft":    playerWithNFT.NFT,
						"name":   nftDetails.Name,
					})
				}
			} else {
				log.Printf("CardPack already minted: %v", cardPack.ID)

			}
		}
	}
	log.Printf("Minted players: %v", mintedPlayers)

	var resultPick *schema.GamePick
	if err = e.Store.Transaction(func(s database.Store) error {
		user, err := s.LockUser(userID)
		if err != nil {
			return err
		}

		if game.SportID.Valid {
			if game.Status == database.GameStatusGameplay {
				if powerUp, err := s.GetSportSubstitutionPowerUp(game.SportID.String); err != nil {
					return err
				} else {
					// ROADMAP: check powerUp.Conditions if needed
					if user.Balance-powerUp.Cost < 0 {
						return ErrNotEnoughBalanceForPowerup
					}
					// Create the transaction to update the Balance properly
					if _, err := s.CreateTransaction(&schema.Transaction{
						ID:         uuid.NewString(),
						UserID:     user.ID,
						Amount:     -powerUp.Cost,
						ObjectType: database.TransactionTypeVirtual,
						MatchID:    null.StringFrom(game.MatchID),
						Quantity:   int64(1),
					}); err != nil {
						return err
					}
				}
			}
		} else {
			return stacktrace.Propagate(err,
				fmt.Sprintf("Substitution failed (game_id: %s - player_id: %s): no sport ID  set for the game", game.ID, player.ID))
		}

		// finalize current pick
		pick.EndedAt = null.TimeFrom(now)
		if err := s.UpdateGamePickEndedAt(pick); err != nil {
			return err
		}

		var assignedPlayerID null.String
		// Filter and find the rarest player
		var matchingPlayers []map[string]interface{}
		for _, mintedPlayer := range mintedPlayers {
			if mintedPlayer["nft"].(*schema.NFTBucket).OptaID == player.ImportID {
				matchingPlayers = append(matchingPlayers, mintedPlayer)
			}
		}

		if len(matchingPlayers) > 0 {
			rarestPlayer := findRarestPlayer(matchingPlayers)
			assignedPlayerID = null.StringFrom(rarestPlayer["player"].(*schema.AssignedPlayer).ID)
		}

		// create new pick
		newPick, err := s.CreateGamePick(&schema.GamePick{
			ID:               uuid.New().String(),
			CreatedAt:        now,
			GameID:           game.ID,
			PlayerID:         player.ID,
			Position:         pick.Position,
			Minute:           game.R.Match.Minute,
			Second:           game.R.Match.Second,
			UserSwapped:      game.Status == database.GameStatusGameplay,
			AssignedPlayerID: assignedPlayerID,
		})
		if err != nil {
			return err
		}
		resultPick = newPick

		return nil
	}); err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToGamePick(resultPick,
		util.IsPlayerPlaying(player.ID, game.R.Match.R.MatchPlayers)))
}

// GetGameEvents godoc
// @Summary Get game events
// @Description Get game events
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Game ID"
// @Param page query string false "Page"
// @Param page_size query string false "Page size"
// @Success 200 {object} model.GameEvent
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /games/{id}/events [get]
func (e *Env) GetGameEvents(c echo.Context) error {
	gameID := c.Param("id")
	ok, err := e.Store.IsGameBelongToUser(gameID, userID(c))
	if err != nil {
		return err
	}

	if !ok {
		return ErrObjectNotFound
	}

	var offset, limit = e.GetOffsetAndLimit(c.QueryParam("page"), c.QueryParam("page_size"))
	events, err := e.Store.GetGameEvents(gameID, offset, limit)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToGameEventSlice(events, e.ActionStore, e.Store))
}

// GetLeaderBoardPosition godoc
// @Summary Get leader board position
// @Description Get leader board position
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Game ID"
// @Success 200 {object} model.LeaderboardPosition
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /games/{id}/leaderboard [get]
func (e *Env) GetLeaderBoardPosition(c echo.Context) error {
	gameID := c.Param("id")
	ok, err := e.Store.IsGameBelongToUser(gameID, userID(c))
	if err != nil {
		return err
	}

	if !ok {
		return ErrObjectNotFound
	}

	lb, err := e.Store.GetLeaderBoardForGame(gameID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToLeaderBoardPosition(lb))
}

// GetGameHistory godoc
// @Summary Get game history
// @Description Get game history
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.HistoricalGame
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Router /games/history [get]
func (e *Env) GetGameHistory(c echo.Context) error {
	userID := userID(c)
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	games, err := e.Store.GetFinishedGames(userID, 20)
	if err != nil {
		return stacktrace.Propagate(err, "Could not get games from database for user %s", userID)
	}

	var result []*model.HistoricalGame
	for _, game := range games {
		if game == nil || game.R == nil || game.R.Match == nil || game.R.Match.R == nil ||
			game.R.Match.R.HomeTeam == nil || game.R.Match.R.AwayTeam == nil {
			return stacktrace.NewError("Missing essential game or match data for game ID: %s", game.ID)
		}

		matchName := fmt.Sprintf("%s vs %s",
			apiconv.GetTeamDisplayName(game.R.Match.R.HomeTeam),
			apiconv.GetTeamDisplayName(game.R.Match.R.AwayTeam),
		)

		lb, err := e.Store.GetLeaderBoardForGame(game.ID)
		if err != nil {
			return stacktrace.Propagate(err, "Could not get leaderboard for game %s", game.ID)
		}
		if lb == nil {
			return stacktrace.NewError("Leaderboard data is nil for game ID: %s", game.ID)
		}

		// Try to find any prizes for this match
		tra, err := e.Store.FindRewardTransaction(game.MatchID, game.UserID)
		var prize float64
		if err != nil {
			if err != sql.ErrNoRows {
				return stacktrace.Propagate(err, "Could not find transaction for game %s", game.ID)
			}
		} else {
			prize = tra.Amount
		}

		// Get total player count for this game
		playerCount, err := e.Store.GetPlayerCount(game.MatchID)
		if err != nil {
			return stacktrace.Propagate(err, "Could not get player count for game %s", game.ID)
		}

		// Constructing the historical game
		historicalGame := apiconv.ToHistoricalGame(
			game.R.Match.ID,
			matchName,
			apiconv.GetTeamDisplayName(game.R.Match.R.HomeTeam),
			apiconv.GetTeamDisplayName(game.R.Match.R.AwayTeam),
			game.R.Match.HomeScore,
			game.R.Match.AwayScore,
			lb.Position.Int,
			int(playerCount),
			game.Num.Int,
			game.Score,
			game.R.Match.MatchTime,
			prize,
			game.R.GamePicks,
			game.R.Match,
		)
		result = append(result, historicalGame)
	}

	return e.RespondSuccess(c, result)
}

// GetUnnotifiedGames godoc
// @Summary Get unnotified games for a user
// @Description Get all games with notified = false for a specific user
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []model.Game
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /games/unnotified [get]
func (e *Env) GetUnnotifiedGames(c echo.Context) error {
	games, err := e.Store.GetUnnotifiedGamesByUserID(userID(c))
	if err != nil {
		fmt.Println("Error fetching unnotified games:", err)
		return err
	}

	if games == nil {
		fmt.Println("No unnotified games found for user:", userID(c))
		return e.RespondSuccess(c, []*model.Game{})
	}

	return e.RespondSuccess(c, apiconv.ToGameSlice(games))
}

// SetGameNotified godoc
// @Summary Set game as notified
// @Description Set the notified field to true for a specific game
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Game ID"
// @Success 200 {object} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /games/{id}/notify [post]
func (e *Env) SetGameNotified(c echo.Context) error {
	gameID := c.Param("id")

	err := e.Store.SetGameNotified(gameID)
	if err != nil {
		// Log the error for internal tracking
		_ = stacktrace.Propagate(err, "Error setting game %s as notified", gameID)
		// Return a detailed error response
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "game not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return e.RespondSuccess(c, "ok")
}

func findRarestPlayer(players []map[string]interface{}) map[string]interface{} {
	rarityOrder := map[string]int{
		"common":     0,
		"uncommon":   1,
		"rare":       2,
		"ultra_rare": 3,
		"legendary":  4,
	}

	rarestPlayer := players[0]
	for _, player := range players {
		playerRarity := player["player"].(*schema.AssignedPlayer).Rarity.String
		rarestPlayerRarity := rarestPlayer["player"].(*schema.AssignedPlayer).Rarity.String

		if rarityOrder[playerRarity] > rarityOrder[rarestPlayerRarity] {
			rarestPlayer = player
		}
	}

	return rarestPlayer
}

// SetAllGamesNotified godoc
// @Summary Set all games as notified
// @Description Set the notified field to true for all games where it is currently false
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Router /games/notify/all [post]
func (e *Env) SetAllGamesNotified(c echo.Context) error {
	// Update all games where notified is false
	err := e.Store.SetAllGamesNotified()
	if err != nil {
		// Log the error for internal tracking
		_ = stacktrace.Propagate(err, "Error setting all games as notified")
		// Return a detailed error response
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return e.RespondSuccess(c, "ok")
}

// JoinGame godoc
// @Summary Join game
// @Description Join game
// @Tags games
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Game ID"
// @Success 200 {object} model.Game
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Router /games/{id}/join [post]
func (e *Env) NewJoinGame(c echo.Context) error {
	log.Println("JoinGame called")
	var in model.NewGameJoinRequest

	// Parse request body
	if err := e.ParseBody(c, &in); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	log.Println("Request body parsed successfully")

	userID := userID(c)
	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		log.Printf("Error retrieving user by ID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user")
	}
	log.Printf("User retrieved: %v", user)

	isBanned := false
	if user.R != nil && user.R.BanPenalties != nil {
		for _, ban := range user.R.BanPenalties {
			if !ban.EndTime.Valid || ban.EndTime.Time.After(time.Now()) {
				isBanned = true
			}
		}
	}
	if isBanned {
		log.Println("User is banned")
		return echo.NewHTTPError(http.StatusForbidden, "User is banned")
	}

	if len(in.Picks) != 4 {
		log.Println("Invalid number of picks")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid number of picks")
	}
	log.Println("Valid number of picks")

	assignedCardPacks, err := e.Store.GetAssignedCardPacksByUserID(userID)
	if err != nil {
		log.Printf("Error retrieving assigned card packs by user ID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve assigned card packs")
	}
	log.Printf("Assigned card packs retrieved: %v", assignedCardPacks)

	match, err := e.Store.GetMatchByID(*in.MatchID)
	if err != nil {
		log.Printf("Error retrieving match by ID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve match")
	}
	log.Printf("Match retrieved: %v", match)

	var playersWithNFT []PlayerWithNFT
	for _, p := range in.Picks {
		// Retrieve player by the pick ID
		player, err := e.Store.GetPlayerByID(*p.Pick)
		if err != nil {
			log.Printf("Error retrieving player by ID: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve player")
		}

		// Append both player and the corresponding nft
		playersWithNFT = append(playersWithNFT, PlayerWithNFT{
			Player: *player,
			NFT:    *p.Nft, // Assuming p.Nft is a pointer, dereference it
		})
	}
	log.Printf("Players with NFTs retrieved: %v", playersWithNFT)

	if match.MatchType == database.MatchTypeUnknown {
		log.Println("Invalid match type")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid match type")
	}

	// Check whether game already exists
	activeGameID, err := e.Store.GetActiveGameIDForMatch(match.ID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Game not exists, create new one
			log.Println("Creating new game")
			err = e.Store.Transaction(func(s database.Store) error {
				// Find proper game status
				var gameStatus = database.GameStatusWaiting
				if match.Status == database.MatchStatusGame {
					gameStatus = database.GameStatusGameplay
				}

				// Calculate game num
				gameNum, err := s.GetNumberOfGames(userID)
				if err != nil {
					log.Printf("Error getting number of games: %v", err)
					return err
				}

				// Create game
				game, err := s.CreateGame(&schema.Game{
					ID:               uuid.New().String(),
					UserID:           userID,
					MatchID:          match.ID,
					Status:           gameStatus,
					Premium:          user.Premium,
					SubscriptionTier: user.SubscriptionTier,
					Num:              null.IntFrom(gameNum + 1),
					SportID:          match.SportID,
					Notified:         false,
				})
				if err != nil {
					log.Printf("Error creating game: %v", err)
					return err
				}

				// Create picks
				for idx, playerWithNFT := range playersWithNFT {
					// Create a game pick object with the pick's related nft as AssignedPlayerID
					gamePick := &schema.GamePick{
						ID:               uuid.New().String(),
						GameID:           game.ID,
						PlayerID:         playerWithNFT.Player.ID, // Access the player ID
						Position:         idx + 1,
						Minute:           match.Minute,
						Second:           match.Second,
						UserSwapped:      true,
						AssignedPlayerID: null.NewString(playerWithNFT.NFT, true), // Convert to null.String
					}

					_, err = s.CreateGamePick(gamePick)
					if err != nil {
						log.Printf("Error creating game pick for player %s: %v", playerWithNFT.Player.ID, err)
						return err
					}
				}

				// Notify that game was created
				_, err = s.CreateAMQPEvent(config.AmqpGamesExchange(), "game_updated", map[string]interface{}{
					"game_id":  game.ID,
					"match_id": match.ID,
				})
				if err != nil {
					log.Printf("Error creating AMQP event: %v", err)
					return err
				}

				activeGameID = game.ID

				return nil
			})
			if err != nil {
				log.Printf("Error creating game: %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create game")
			}
		} else {
			log.Printf("Error retrieving active game: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve active game")
		}
	}

	// Game exists, query from db
	game, err := e.Store.GetGameByID(activeGameID, userID)
	if err != nil {
		log.Printf("Error retrieving game by ID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve game")
	}
	log.Printf("Game retrieved: %v", game)

	return e.RespondSuccess(c, apiconv.ToGame(game))
}
