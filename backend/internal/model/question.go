package model

import "time"

const (
	QuestionStatusOpen     = "open"
	QuestionStatusResolved = "resolved"
	QuestionStatusClosed   = "closed"
)

// Question 对应 questions 表。
type Question struct {
	ID        int64  `json:"id"`
	AuthorID  int64  `json:"author_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	ViewCount int64  `json:"view_count"`
	Base
	SoftDelete

	Author       User          `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
	QuestionTags []QuestionTag `json:"question_tags,omitempty" gorm:"foreignKey:QuestionID"`
	Answers      []Answer      `json:"answers,omitempty" gorm:"foreignKey:QuestionID"`
}

func (Question) TableName() string {
	return "questions"
}

// QuestionTag 对应 question_tags 表，移除标签时保留软删除历史。
type QuestionTag struct {
	ID         int64 `json:"id"`
	QuestionID int64 `json:"question_id"`
	TagID      int64 `json:"tag_id"`
	Base
	SoftDelete

	Question Question `json:"-" gorm:"foreignKey:QuestionID"`
	Tag      Tag      `json:"tag" gorm:"foreignKey:TagID"`
}

func (QuestionTag) TableName() string {
	return "question_tags"
}

// Answer 对应 answers 表。
type Answer struct {
	ID         int64      `json:"id"`
	QuestionID int64      `json:"question_id"`
	AuthorID   int64      `json:"author_id"`
	Content    string     `json:"content"`
	IsAccepted bool       `json:"is_accepted"`
	AcceptedAt *time.Time `json:"accepted_at"`
	Base
	SoftDelete

	Question    Question     `json:"-" gorm:"foreignKey:QuestionID"`
	Author      User         `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
	AnswerLikes []AnswerLike `json:"answer_likes,omitempty" gorm:"foreignKey:AnswerID"`
}

func (Answer) TableName() string {
	return "answers"
}

// AnswerLike 对应 answer_likes 表，取消点赞时保留软删除历史。
type AnswerLike struct {
	ID       int64 `json:"id"`
	UserID   int64 `json:"user_id"`
	AnswerID int64 `json:"answer_id"`
	Base
	SoftDelete

	User   User   `json:"-" gorm:"foreignKey:UserID"`
	Answer Answer `json:"answer,omitempty" gorm:"foreignKey:AnswerID"`
}

func (AnswerLike) TableName() string {
	return "answer_likes"
}
