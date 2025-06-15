package dbstore

import (
	"context"
	"fmt"

	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/types"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// TODO: get users, rooms from db as name and send the object to the client in the response

// CreateChatMessage inserts a new ChatMessage record into the database.
func (s *DBStore) CreateChatMessage(chatMessage *schema.ChatMessage) error {
	if chatMessage.ID == "" {
		chatMessage.ID = uuid.New().String()
	}
	return chatMessage.Insert(s.db, boil.Infer())
}

// GetChatRoomByName retrieves a ChatRoom record by its name.
func (s *DBStore) GetChatRoomByName(name string) (*schema.ChatRoom, error) {
	return schema.ChatRooms(schema.ChatRoomWhere.Name.EQ(name)).One(s.db)
}

// GetChatMessagesByRoomID retrieves all ChatMessage records for a given ChatRoom ID.
func (s *DBStore) GetChatMessagesByRoomID(roomID string, offset, limit, minutesPrior int) (*[]types.ChatMessage, error) {
	var msgs []types.ChatMessage

	var q = schema.NewQuery(
		qm.Select("cm.id, cm.message, cm.created_at, cm.updated_at, u.name as user_name, u.avatar_url"),
		qm.From(schema.TableNames.ChatMessages+" cm"),
		qm.InnerJoin(schema.TableNames.Users+" u on u.id = sender_id"),
		// schema.ChatMessageWhere.RoomID.EQ(roomID),
		qm.Where("cm.room_id = ?", roomID),
		qm.Where(fmt.Sprintf("cm.created_at >= NOW() - interval '%d minutes'", minutesPrior)),
		qm.OrderBy("cm.created_at DESC"),
		qm.Limit(limit),
		qm.Offset(offset),
	// ).All(s.db)
	)
	err := q.Bind(context.Background(), s.db, &msgs)

	if msgs == nil {
		msgs = make([]types.ChatMessage, 0)
	}

	return &msgs, err
}

// DeleteChatRoomMesageByID deletes a specific (by id) message from a specific room
func (s *DBStore) DeleteChatRoomMessageByID(roomID, messageID string) (bool, error) {
	// schema.ChatMessage().delete
	if c, err := schema.ChatMessages(schema.ChatMessageWhere.RoomID.EQ(roomID), schema.ChatMessageWhere.ID.EQ(messageID)).DeleteAll(s.db); err != nil {
		return false, err
	} else {
		return c > 0, nil
	}
}

// GetChatRoomMember gets the information of a specific user in a chatroom
func (s *DBStore) GetChatRoomMember(roomID, userID string) (*schema.ChatRoomMember, error) {
	return schema.ChatRoomMembers(
		schema.ChatRoomMemberWhere.RoomID.EQ(roomID),
		schema.ChatRoomMemberWhere.UserID.EQ(userID)).One(s.db)
}

// GetChatRoomMessageByID gets the information of a specific message in a chatroom
func (s *DBStore) GetChatRoomMessageByID(roomID, messageID string) (*schema.ChatMessage, error) {
	return schema.ChatMessages(
		schema.ChatMessageWhere.RoomID.EQ(roomID),
		schema.ChatMessageWhere.ID.EQ(messageID)).One(s.db)
}

// GetChatRoomByMatchID gets the information of a specific match chatroom
func (s *DBStore) GetChatRoomByMatchID(matchID string) (*schema.ChatRoom, error) {
	return schema.ChatRooms(
		schema.ChatRoomWhere.MatchID.EQ(null.StringFrom(matchID))).One(s.db)
}
