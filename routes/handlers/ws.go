package handlers

import (
	"fmt"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/util"
	"github.com/labstack/echo/v4"
	"gopkg.in/olahol/melody.v1"
)

const (
	WsLeaderboardConfigName = "leaderboard"
	WsParamMatchID          = "match_id"
	WsEventPipeName         = "game"
)

func (e *Env) EventPipe(c echo.Context) error {
	config := e.WebsocketConfigs[WsEventPipeName]
	token := c.QueryParam(WsParamToken)
	println(token)
	return config.Manager.HandleRequestWithKeys(c.Response().Writer, c.Request(), map[string]interface{}{
		WsParamToken: c.QueryParam(WsParamToken),
	})
}

func (e *Env) InitEventPipe() {
	config := e.WebsocketConfigs[WsEventPipeName]

	config.Manager.HandleConnect(func(s *melody.Session) {
		userID, exists := s.Get("user_id")
		println(userID, exists)
		if !exists {
			println("not exists user")
			qToken, existsToken := s.Get(WsParamToken)
			if !existsToken || qToken == nil {
				_ = s.CloseWithMsg(melody.FormatCloseMessage(melody.ClosePolicyViolation, ErrJwtUnauthorized))
				return
			}

			// Ensure token is a string
			tokenString, ok := qToken.(string)
			if !ok || tokenString == "" {
				_ = s.CloseWithMsg(melody.FormatCloseMessage(melody.ClosePolicyViolation, ErrJwtUnauthorized))
				return
			}

			// Parse JWT token
			_, claims, err := util.ParseJWTWithClaims(JwtSecretKey, tokenString)
			if err != nil {
				_ = s.CloseWithMsg(melody.FormatCloseMessage(melody.ClosePolicyViolation, err.Error()))
				return
			}

			// Set user_id in session
			userID = claims[ClaimsUserID]
			s.Set("user_id", userID)

		}

		// Ensure user ID is a string
		userIDString, ok := userID.(string)
		if !ok {
			_ = s.CloseWithMsg(melody.FormatCloseMessage(melody.ClosePolicyViolation, ErrJwtUnauthorized))
			return
		}

		config.EventBus.AddListener(userIDString, s)
	})

	config.Manager.HandleDisconnect(func(s *melody.Session) {
		userID, exists := s.Get("user_id")
		if !exists {
			return
		}
		userIDString := userID.(string)

		config.EventBus.RemoveListener(userIDString, s)
	})
}

func (e *Env) LeaderboardWebsocketHandler(c echo.Context) error {
	e.initLeaderboardPipe(WsLeaderboardConfigName)
	ws := e.WebsocketConfigs[WsLeaderboardConfigName].Manager
	return ws.HandleRequestWithKeys(c.Response().Writer, c.Request(), map[string]interface{}{
		WsParamMatchID: c.QueryParam(WsParamMatchID),
	})
}

func (e *Env) initLeaderboardPipe(wsConfigName string) {
	ws := e.WebsocketConfigs[wsConfigName]
	ws.Manager.HandleConnect(func(s *melody.Session) {
		qMatchID, exists := s.Get(WsParamMatchID)
		if matchIDString, ok := qMatchID.(string); !ok {
			_ = s.CloseWithMsg(melody.FormatCloseMessage(melody.CloseUnsupportedData, ErrJwtUnauthorized))
		} else {
			if !exists || matchIDString == "" {
				_ = s.CloseWithMsg(melody.FormatCloseMessage(melody.CloseUnsupportedData, ErrJwtUnauthorized))
				return
			}
			_, err := e.Store.GetMatchByID(matchIDString)
			if err != nil {
				_ = s.CloseWithMsg(melody.FormatCloseMessage(melody.CloseUnsupportedData, err.Error()))
				return
			}

			ws.EventBus.AddListener(matchIDString, s)

			fmt.Printf("Connected to WS Leaderboard for match: %s", matchIDString)
		}
	})

	ws.Manager.HandleDisconnect(func(s *melody.Session) {
		qMatchID, exists := s.Get(WsParamMatchID)
		if !exists {
			return
		}
		matchIDString := qMatchID.(string)

		ws.EventBus.RemoveListener(matchIDString, s)

		fmt.Printf("Disconnected to WS Leaderboard for match: %s", matchIDString)
	})
}
