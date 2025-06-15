package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/itemslabs/clubz-api/types"
	"github.com/itemslabs/clubz-api/util"
	"gopkg.in/olahol/melody.v1"

	"github.com/labstack/echo/v4"
)

var (
	ErrChatRoomInvalid     = &APIError{Message: "Invalid chat room", StatusCode: http.StatusNotFound}
	ErrChatRoomNotFound    = &APIError{Message: "Chat room not found", StatusCode: http.StatusNotFound, Code: http.StatusNotFound}
	ErrChatMessageTooShort = &APIError{Message: "Chat message too short"}
	ErrChatMessageInvalid  = &APIError{Message: "Chat message not valid"}
)

const (
	ErrJwtInvalidToken = "Invalid token"
	ErrJwtUnauthorized = "Unauthorized"
	WsChatConfigname   = "chat"
	WsParamToken       = "token"
	WsParamRoom        = "room"
	ClaimsUserID       = "user_id"
)

func (e *Env) ChatWebsocketHandler(c echo.Context) error {
	e.initChatPipe(WsChatConfigname)
	ws := e.WebsocketConfigs[WsChatConfigname].Manager
	return ws.HandleRequestWithKeys(c.Response().Writer, c.Request(), map[string]interface{}{
		WsParamRoom:  c.QueryParam(WsParamRoom),
		WsParamToken: c.QueryParam(WsParamToken),
	})
}

func (e *Env) initChatPipe(wsConfigName string) {
	ws := e.WebsocketConfigs[wsConfigName].Manager
	ws.HandleConnect(func(s *melody.Session) {
		qToken, exists := s.Get(WsParamToken)
		if tokenString, ok := qToken.(string); !ok {
			_ = s.CloseWithMsg(melody.FormatCloseMessage(melody.ClosePolicyViolation, ErrJwtUnauthorized))
		} else {
			if !exists || tokenString == "" {
				_ = s.CloseWithMsg(melody.FormatCloseMessage(melody.ClosePolicyViolation, ErrJwtUnauthorized))
				return
			}
			_, claims, err := util.ParseJWTWithClaims(JwtSecretKey, tokenString)
			if err != nil {
				_ = s.CloseWithMsg(melody.FormatCloseMessage(melody.ClosePolicyViolation, err.Error()))
				return
			}
			fmt.Printf("userID %s connected to chat\n", claims[ClaimsUserID])
		}
	})

	ws.HandleDisconnect(func(s *melody.Session) {
		// A connected user would always have a valid token in the querystring
		qToken, _ := s.Get(WsParamToken)
		if tokenString, ok := qToken.(string); ok {
			_, claims, err := util.ParseJWTWithClaims(JwtSecretKey, tokenString)
			if err != nil {
				_ = s.CloseWithMsg(melody.FormatCloseMessage(melody.ClosePolicyViolation, err.Error()))
				return
			}
			fmt.Printf("userID %s disconnected from chat\n", claims[ClaimsUserID])
		}
	})
}

// GetChatMessages godoc
// @Summary Get chat messages
// @Description Get chat messages
// @Tags chat
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param room path string true "Room name"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} *schema.ChatMessageSlice
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Router /chat/{room} [get]
func (e *Env) GetChatMessages(c echo.Context) error {
	roomID := c.Param("room")
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
	offset, limit := e.GetOffsetAndLimit(c.QueryParam("page"), c.QueryParam("page_size"))
	strMinutesPrior := c.QueryParam("minutes_prior")
	minutesPrior, err := strconv.Atoi(strMinutesPrior)
	if err != nil {
		minutesPrior = 120
	}
	fmt.Print("minutesPrior: ", minutesPrior)
	msgs, err := e.Store.GetChatMessagesByRoomID(roomID, offset, limit, minutesPrior)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, msgs)
}

// PostChatMessage godoc
// @Summary Post chat message
// @Description Post chat message
// @Tags chat
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 204 {string} string	"ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Router /chat [post]
func (e *Env) PostChatMessage(c echo.Context) error {
	userID := userID(c)
	room := c.Param("room")
	if room == "" {
		return ErrChatRoomInvalid
	}
	if !ValidateUUID(room) {
		if chatroom, err := e.Store.GetChatRoomByName(room); err == nil {
			room = chatroom.ID
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
		RoomID:   room,
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
	})
	_ = ws.Broadcast(b)

	return e.RespondNoContent(c)
}

func (e *Env) DeleteChatRoomMessage(c echo.Context) error {
	userID := userID(c)
	roomID := c.Param("room")
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
