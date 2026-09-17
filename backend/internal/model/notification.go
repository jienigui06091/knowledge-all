package model

import "time"

const (
	NotificationTypeAnswerToQuestion   = "answer_to_question"
	NotificationTypeAnswerLiked        = "answer_liked"
	NotificationTypeKnowledgeFavorited = "knowledge_favorited"
	NotificationTypeAnswerAccepted     = "answer_accepted"
	NotificationTypeChatMessage        = "chat_message"
)

// Notification 对应 notifications 表。
type Notification struct {
	ID               int64      `json:"id"`
	RecipientID      int64      `json:"recipient_id"`
	ActorID          *int64     `json:"actor_id"`
	NotificationType string     `json:"notification_type"`
	ResourceType     string     `json:"resource_type"`
	ResourceID       *int64     `json:"resource_id"`
	Title            string     `json:"title"`
	Content          string     `json:"content"`
	Payload          JSONB      `json:"payload" gorm:"type:jsonb"`
	IsRead           bool       `json:"is_read"`
	ReadAt           *time.Time `json:"read_at"`
	CreatedAt        time.Time  `json:"created_at"`
	SoftDelete

	Recipient User  `json:"recipient,omitempty" gorm:"foreignKey:RecipientID"`
	Actor     *User `json:"actor,omitempty" gorm:"foreignKey:ActorID"`
}

func (Notification) TableName() string {
	return "notifications"
}
