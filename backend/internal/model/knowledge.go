package model

import "time"

const (
	CategoryStatusActive   = "active"
	CategoryStatusDisabled = "disabled"

	KnowledgeStatusDraft     = "draft"
	KnowledgeStatusPublished = "published"
	KnowledgeStatusArchived  = "archived"
)

// Category 对应 categories 表。
type Category struct {
	ID          int64   `json:"id"`
	ParentID    *int64  `json:"parent_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	SortOrder   int     `json:"sort_order"`
	Status      string  `json:"status"`
	CreatedBy   *int64  `json:"created_by"`
	Base
	SoftDelete

	Parent   *Category  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Creator  *User      `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
	Children []Category `json:"children,omitempty" gorm:"foreignKey:ParentID"`
}

func (Category) TableName() string {
	return "categories"
}

// Tag 对应 tags 表。
type Tag struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Color     *string `json:"color"`
	CreatedBy *int64  `json:"created_by"`
	Base
	SoftDelete

	Creator *User `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
}

func (Tag) TableName() string {
	return "tags"
}

// Knowledge 对应 knowledge 表。
type Knowledge struct {
	ID          int64      `json:"id"`
	AuthorID    int64      `json:"author_id"`
	CategoryID  *int64     `json:"category_id"`
	Title       string     `json:"title"`
	Summary     *string    `json:"summary"`
	Content     string     `json:"content"`
	Status      string     `json:"status"`
	ViewCount   int64      `json:"view_count"`
	PublishedAt *time.Time `json:"published_at"`
	Base
	SoftDelete

	Author        User           `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
	Category      *Category      `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	KnowledgeTags []KnowledgeTag `json:"knowledge_tags,omitempty" gorm:"foreignKey:KnowledgeID"`
}

func (Knowledge) TableName() string {
	return "knowledge"
}

// KnowledgeTag 对应 knowledge_tags 表，移除标签时保留软删除历史。
type KnowledgeTag struct {
	ID          int64 `json:"id"`
	KnowledgeID int64 `json:"knowledge_id"`
	TagID       int64 `json:"tag_id"`
	Base
	SoftDelete

	Knowledge Knowledge `json:"-" gorm:"foreignKey:KnowledgeID"`
	Tag       Tag       `json:"tag" gorm:"foreignKey:TagID"`
}

func (KnowledgeTag) TableName() string {
	return "knowledge_tags"
}

// Favorite 对应 favorites 表，取消收藏时保留软删除历史。
type Favorite struct {
	ID          int64 `json:"id"`
	UserID      int64 `json:"user_id"`
	KnowledgeID int64 `json:"knowledge_id"`
	Base
	SoftDelete

	User      User      `json:"-" gorm:"foreignKey:UserID"`
	Knowledge Knowledge `json:"knowledge,omitempty" gorm:"foreignKey:KnowledgeID"`
}

func (Favorite) TableName() string {
	return "favorites"
}
