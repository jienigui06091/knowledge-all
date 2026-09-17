package model

import "time"

const (
	ConversationTypePrivate = "private"
	ConversationTypeGroup   = "group"

	ConversationMemberRoleOwner  = "owner"
	ConversationMemberRoleMember = "member"

	MembershipStatusActive  = "active"
	MembershipStatusLeft    = "left"
	MembershipStatusRemoved = "removed"

	MessageTypeText   = "text"
	MessageTypeImage  = "image"
	MessageTypeFile   = "file"
	MessageTypeSystem = "system"
)

// Conversation 对应 conversations 表。
type Conversation struct {
	ID               int64      `json:"id"`
	ConversationType string     `json:"conversation_type"`
	Name             *string    `json:"name"`
	CreatorID        int64      `json:"creator_id"`
	DirectUserLowID  *int64     `json:"direct_user_low_id"`
	DirectUserHighID *int64     `json:"direct_user_high_id"`
	LastMessageID    *int64     `json:"last_message_id"`
	LastMessageAt    *time.Time `json:"last_message_at"`
	Base
	SoftDelete

	Creator     User                 `json:"creator,omitempty" gorm:"foreignKey:CreatorID"`
	LastMessage *Message             `json:"last_message,omitempty" gorm:"foreignKey:LastMessageID"`
	Members     []ConversationMember `json:"members,omitempty" gorm:"foreignKey:ConversationID"`
	Messages    []Message            `json:"messages,omitempty" gorm:"foreignKey:ConversationID"`
}

func (Conversation) TableName() string {
	return "conversations"
}

// ConversationMember 对应 conversation_members 表，以复合主键保存成员历史。
type ConversationMember struct {
	ConversationID    int64      `json:"conversation_id" gorm:"primaryKey"`
	UserID            int64      `json:"user_id" gorm:"primaryKey"`
	MemberRole        string     `json:"member_role"`
	MembershipStatus  string     `json:"membership_status"`
	JoinedAt          time.Time  `json:"joined_at"`
	LeftAt            *time.Time `json:"left_at"`
	LastReadMessageID *int64     `json:"last_read_message_id"`
	IsPinned          bool       `json:"is_pinned"`
	IsMuted           bool       `json:"is_muted"`

	Conversation    Conversation `json:"-" gorm:"foreignKey:ConversationID"`
	User            User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	LastReadMessage *Message     `json:"last_read_message,omitempty" gorm:"foreignKey:LastReadMessageID"`
}

func (ConversationMember) TableName() string {
	return "conversation_members"
}

// Message 对应 messages 表。
type Message struct {
	ID               int64      `json:"id"`
	ConversationID   int64      `json:"conversation_id"`
	SenderID         int64      `json:"sender_id"`
	ReplyToMessageID *int64     `json:"reply_to_message_id"`
	ClientMessageID  string     `json:"client_message_id"`
	MessageType      string     `json:"message_type"`
	Content          string     `json:"content"`
	IsRecalled       bool       `json:"is_recalled"`
	RecalledAt       *time.Time `json:"recalled_at"`
	Base
	SoftDelete

	Conversation Conversation `json:"-" gorm:"foreignKey:ConversationID"`
	Sender       User         `json:"sender,omitempty" gorm:"foreignKey:SenderID"`
	ReplyTo      *Message     `json:"reply_to,omitempty" gorm:"foreignKey:ReplyToMessageID"`
}

func (Message) TableName() string {
	return "messages"
}
