package model

import (
	"sync"
	"testing"

	"gorm.io/gorm/schema"
)

func TestModelsMatchSchemaTables(t *testing.T) {
	testCases := []struct {
		name       string
		value      any
		tableName  string
		softDelete bool
	}{
		{"role", Role{}, "roles", false},
		{"user", User{}, "users", true},
		{"user role", UserRole{}, "user_roles", true},
		{"category", Category{}, "categories", true},
		{"tag", Tag{}, "tags", true},
		{"knowledge", Knowledge{}, "knowledge", true},
		{"knowledge tag", KnowledgeTag{}, "knowledge_tags", true},
		{"favorite", Favorite{}, "favorites", true},
		{"question", Question{}, "questions", true},
		{"question tag", QuestionTag{}, "question_tags", true},
		{"answer", Answer{}, "answers", true},
		{"answer like", AnswerLike{}, "answer_likes", true},
		{"conversation", Conversation{}, "conversations", true},
		{"conversation member", ConversationMember{}, "conversation_members", false},
		{"message", Message{}, "messages", true},
		{"notification", Notification{}, "notifications", true},
	}

	cache := &sync.Map{}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parsed, err := schema.Parse(testCase.value, cache, schema.NamingStrategy{})
			if err != nil {
				t.Fatalf("parse model: %v", err)
			}
			if parsed.Table != testCase.tableName {
				t.Fatalf("table name = %q, want %q", parsed.Table, testCase.tableName)
			}

			_, hasDeletedAt := parsed.FieldsByDBName["deleted_at"]
			if hasDeletedAt != testCase.softDelete {
				t.Fatalf("deleted_at mapping = %t, want %t", hasDeletedAt, testCase.softDelete)
			}
		})
	}
}
