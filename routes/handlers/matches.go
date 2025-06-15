package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/itemslabs/clubz-api/config"
	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/apiconv"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/itemslabs/clubz-api/types"
	"github.com/labstack/echo/v4"
)

type MatchService struct {
	// store schema.MatchSlice
}

// UpcomingMatches godoc
// @Summary Get upcoming matches
// @Description Get upcoming matches
// @ID upcoming-matches
// @Produce json
// @Success 200 {object} schema.MatchSlice
// @Router /matches [get]
func (e *Env) UpcomingMatches(c echo.Context) error {
	now := time.Now()
	from := now.Add(-config.MatchDisplayLowDelta())
	// add 168 hours to the current time
	to := now.Add(168 * time.Hour)

	if c.QueryParam("debug48HoursAgo") != "" {
		from = now.Add(-48 * time.Hour)
	}

	if c.QueryParam("start") != "" && c.QueryParam("end") != "" {
		if start, err := time.Parse("2006-01-02", c.QueryParam("start")); err != nil {
			return err
		} else {
			from = start
		}
		if end, err := time.Parse("2006-01-02", c.QueryParam("end")); err != nil {
			return err
		} else {
			to = end
		}
	}

	var matches schema.MatchSlice
	if c.QueryParam("week") != "" { // pivot to define if we are searching range of dates or game week
		strGW := c.QueryParam("week")
		if gw, err := strconv.Atoi(strGW); err != nil {
			return err
		} else {
			if gwMatches, err := e.Store.GetMatchesInGameWeek(gw); err != nil {
				return err
			} else {
				matches = gwMatches
			}
		}
	} else {
		if pMatches, err := e.Store.GetMatchesInPeriod(from, to); err != nil {
			return err
		} else {
			matches = pMatches
		}
	}

	// calculate number of future matches
	var futureMatches = 0
	for _, m := range matches {
		if m.MatchTime.After(now) {
			futureMatches += 1
		}
	}
	// if no future matches, take next active one
	if futureMatches == 0 {
		match, err := e.Store.GetNextActiveMatch(now)
		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}

			// in case of no rows do nothing
		} else {
			matches = append(matches, match)
		}
	}

	// Fetch player counts for all matches and store them in a map
	playerCounts := make(map[string]int64)
	for _, m := range matches {
		playerCount, err := e.Store.GetPlayerCount(m.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve player count")
		}
		playerCounts[m.ID] = int64(playerCount)
	}

	return e.RespondSuccess(c, apiconv.ToMatchSliceWithPlayerCount(matches, playerCounts))

}

// GetMatchInfo godoc
// @Summary Get match info
// @Description Get match info
// @ID match-info
// @Produce json
// @Param id path string true "Match ID"
// @Success 200 {object} schema.Match
// @Router /matches/{id} [get]
func (e *Env) GetMatchInfo(c echo.Context) error {
	matchID := c.Param("id")

	// Retrieve match information by ID
	match, err := e.Store.GetMatchByID(matchID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Match not found")
	}

	// Get the player count for the match
	playerCount, err := e.Store.GetPlayerCount(matchID)
	if err != nil {

		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve player count")
	}

	// Convert schema match to API model match
	result := apiconv.ToMatch(match)
	result.PlayerCount = int64(playerCount) // Convert int to int64 and assign

	return e.RespondSuccess(c, result)
}

// GetMatchSquad godoc
// @Summary Get match squad
// @Description Get match squad
// @ID match-squad
// @Produce json
// @Param id path string true "Match ID"
// @Success 200 {object} schema.MatchPlayer
// @Router /matches/{id}/squad [get]
func (e *Env) GetMatchSquad(c echo.Context) error {
	players, err := e.Store.GetMatchPlayers(c.Param("id"))
	if err != nil {
		return err
	}
	mpPPGs, err := e.Store.GetMatchPlayersPPG(c.Param("id"))
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToMatchPlayerSlice(players, mpPPGs))
}

// GetLobbyHeadlines godoc
// @Summary Get match headlines
// @Description Get match headlines
// @ID match-headlines
// @Produce json
// @Param id path string true "Match ID"
// @Success 200 {object} schema.MatchHeadlineSlice
// @Router /matches/{id}/headlines [get]
func (e *Env) GetLobbyHeadlines(c echo.Context) error {
	headlines, err := e.Store.GetMatchHeadlines(c.Param("id"), database.MatchHeadlineScreenTypeLobby)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToMatchHeadlineSlice(headlines))
}

// GetGamePlayHeadlines godoc
// @Summary Get match gameplay headlines
// @Description Get match gameplay headlines
// @ID match-gameplay-headlines
// @Produce json
// @Param id path string true "Match ID"
// @Success 200 {object} schema.MatchHeadlineSlice
// @Router /matches/{id}/gameplay/headlines [get]
func (e *Env) GetGamePlayHeadlines(c echo.Context) error {
	headlines, err := e.Store.GetMatchHeadlines(c.Param("id"), database.MatchHeadlineScreenTypeGamePlay)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToMatchHeadlineSlice(headlines))
}

// GetFullTimeHeadlines godoc
// @Summary Get match full time headlines
// @Description Get match full time headlines
// @ID match-full-time-headlines
// @Produce json
// @Param id path string true "Match ID"
// @Success 200 {object} schema.MatchHeadlineSlice
// @Router /matches/{id}/fulltime/headlines [get]
func (e *Env) GetFullTimeHeadlines(c echo.Context) error {
	headlines, err := e.Store.GetMatchHeadlines(c.Param("id"), database.MatchHeadlineScreenTypeFullTime)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToMatchHeadlineSlice(headlines))
}

// FollowingLeaderboard godoc
// @Summary Get following leaderboard
// @Description Get following leaderboard
// @ID following-leaderboard
// @Produce json
// @Param id path string true "Match ID"
// @Success 200 {object} model.LeaderboardEntry
// @Router /matches/{id}/leaderboard/following [get]
func (e *Env) FollowingLeaderboard(c echo.Context) error {
	user, err := e.Store.GetUserByID(userID(c))
	if err != nil {
		return err
	}

	// allow to use this leaderboard only to moderators
	if !user.Moderator {
		return ErrNotAuthorized
	}

	records, err := e.Store.GetFollowingLeaderboard(c.Param("id"), user.ID, 100)
	if err != nil {
		return err
	}

	result := make([]*model.LeaderboardEntry, 0, len(records))
	for _, rec := range records {
		result = append(result, apiconv.ToLeaderBoardEntry(rec))
	}

	return e.RespondSuccess(c, result)
}

// GetLeaderboard
// @Summary Get top N users by score
// @Description Retrieve top N users from the leaderboard by score
// @ID GetLeaderboard
// @Param matchID path string true "Match ID"
// @Param count path int true "Number of users to retrieve"
// @Success 200 {object} []model.LeaderboardEntry "Leaderboard data"
// @Failure 400 {string} string "Invalid count parameter"
// @Failure 500 {string} string "Failed to retrieve leaderboard data"
// @Router /{matchID}/leaderboard/{count} [get]
func (e *Env) GetLeaderboard(c echo.Context) error {
	matchID := c.Param("matchID")
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid count parameter")
	}

	match, err := e.Store.GetMatchByID(matchID)
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("Match: %s not found", matchID))
	}
	if match.Status == database.MatchStatusEnded { //status ended
		records, err := e.Store.GetLeaderboard(matchID, 100)
		if err != nil {
			return err
		}

		result := make([]*model.LeaderboardEntry, 0, len(records))
		for _, rec := range records {
			// spew.Dump(rec)
			// spew.Dump(apiconv.ToLeaderBoardEntry(rec))
			result = append(result, apiconv.ToLeaderBoardEntry(rec))
		}

		return e.RespondSuccess(c, map[string]interface{}{
			"leaderboard": result,
		})
	} else {
		IDsFromRedis, err := e.RedisClient.ZRevRangeByScore(c.Request().Context(), "match:"+matchID+":scores", &redis.ZRangeBy{
			Min:    "-inf",
			Max:    "+inf",
			Offset: 0,
			Count:  int64(count),
		}).Result()

		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to retrieve leaderboard data")
		}

		leaderboard := make([]*model.LeaderboardEntry, 0, len(IDsFromRedis))

		for _, id := range IDsFromRedis {
			dataFromRedis, err := e.RedisClient.HGetAll(c.Request().Context(), "match:"+matchID+":"+id).Result()
			if err != nil {
				fmt.Println("Error fetching data from Redis:", err)
				continue
			}

			flScore, _ := strconv.ParseFloat(dataFromRedis["score"], 64)

			position, _ := strconv.ParseInt(dataFromRedis["position"], 10, 64)
			score := apiconv.ToFloatWithZero(flScore)

			entry := model.LeaderboardEntry{
				Position:      position,
				UserName:      dataFromRedis["user_name"],
				UserID:        dataFromRedis["user_id"],
				UserAvatarURL: dataFromRedis["user_avatar_url"],
				Score:         score,
			}

			leaderboard = append(leaderboard, &entry)
		}
		//var totalPlayers int
		//totalCountKey := fmt.Sprintf("match:%s:fans_playing", matchID)
		//totalPlayers, err = e.RedisClient.Get(c.Request().Context(), totalCountKey).Int()
		//if err != nil {
		//	logrus.WithError(err).Error(fmt.Sprintf("Failed to retrieve leaderboard data: %s", err))
		//}
		return e.RespondSuccess(c, map[string]interface{}{
			"leaderboard": leaderboard,
		})
	}
}

func (e *Env) getMatchRoomID(matchID string) (string, error) {
	var roomID string
	if room, err := e.Store.GetChatRoomByMatchID(matchID); err != nil {
		if err == sql.ErrNoRows {
			defaultRoom, err := e.Store.GetChatRoomByName(config.DefaultChatRoomChannel())
			if err != nil {
				return "", ErrChatRoomNotFound
			}
			roomID = defaultRoom.ID
		} else {
			return "", err
		}
	} else {
		roomID = room.ID
	}
	return roomID, nil
}

func (e *Env) GetMatchChatMessages(c echo.Context) error {
	roomID, err := e.getMatchRoomID(c.Param("id"))
	if err != nil {
		return err
	}
	offset, limit := e.GetOffsetAndLimit(c.QueryParam("page"), c.QueryParam("page_size"))
	strMinutesPrior := c.QueryParam("minutes_prior")
	if !ValidateUUID(roomID) {
		if chatroom, err := e.Store.GetChatRoomByName(roomID); err == nil {
			roomID = chatroom.ID
		} else {
			return ErrChatRoomNotFound
		}
	}
	minutesPrior, err := strconv.Atoi(strMinutesPrior)
	if err != nil {
		minutesPrior = 120
	}
	msgs, err := e.Store.GetChatMessagesByRoomID(roomID, offset, limit, minutesPrior)
	if err != nil {
		return err
	}
	return e.RespondSuccess(c, msgs)
}

func (e *Env) PostMatchChatMessages(c echo.Context) error {
	matchID := c.Param("id")
	roomID, err := e.getMatchRoomID(matchID)
	if err != nil {
		return err
	}
	userID := userID(c)
	println(userID)
	if !ValidateUUID(roomID) {
		if chatroom, err := e.Store.GetChatRoomByName(roomID); err == nil {
			roomID = chatroom.ID
		} else {
			return ErrChatRoomNotFound
		}
	}
	var in model.PostChatMessageRequest
	if err := c.Bind(&in); err != nil {
		return err
	}
	if in.Message == "" {
		return ErrChatMessageTooShort
	}
	dbMsg := &schema.ChatMessage{
		RoomID:   roomID,
		SenderID: userID,
		Message:  in.Message,
	}
	if err := e.Store.CreateChatMessage(dbMsg); err != nil {
		return err
	}
	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		return err
	}

	ws := e.WebsocketConfigs[WsChatConfigname].Manager
	b, _ := json.Marshal(types.ChatMessage{
		ID:        dbMsg.ID,
		Message:   dbMsg.Message,
		CreatedAt: dbMsg.CreatedAt,
		UpdatedAt: dbMsg.UpdatedAt,
		UserName:  user.Name,
		AvatarURL: user.AvatarURL,
		RoomID:    roomID,
		SenderID:  user.ID,
		MatchID:   matchID,
	})
	_ = ws.Broadcast(b)

	return e.RespondNoContent(c)
}

func (e *Env) DeleteMatchChatRoomMessage(c echo.Context) error {
	roomID, err := e.getMatchRoomID(c.Param("id"))
	if err != nil {
		return err
	}
	userID := userID(c)
	messageID := c.Param("messageId")
	if roomID == "" {
		return ErrChatRoomInvalid
	}
	if !ValidateUUID(roomID) {
		if chatroom, err := e.Store.GetChatRoomByName(roomID); err == nil {
			roomID = chatroom.ID
		} else {
			return ErrChatRoomNotFound
		}
	}
	if msg, err := e.Store.GetChatRoomMessageByID(roomID, messageID); err != nil {
		return ErrChatMessageInvalid
	} else if crMember, err := e.Store.GetChatRoomMember(roomID, userID); err != nil {
		return err
	} else if !crMember.Banned && !crMember.Muted && crMember.UserID == msg.SenderID || crMember.Mod {
		if deleted, err := e.Store.DeleteChatRoomMessageByID(roomID, msg.ID); err != nil {
			return err
		} else if deleted {
			return e.RespondNoContent(c)
		}
	}

	return e.RespondNotFound(c)
}
