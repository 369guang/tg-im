package model

import "time"

type ConversationType int

const (
	UserConversation  ConversationType = 1
	GroupConversation ConversationType = 2
)

type Conversation struct {
	ID        uint             `gorm:"primarykey"`        // ID
	Type      ConversationType `gorm:"index" json:"type"` // 会话类型，1：用户，2：群组
	UserID    uint             `json:"user_id"`           // 用户ID
	GroupID   uint             `json:"group_id"`          // 群组ID
	CreatedAt time.Time        // 创建时间（由GORM自动管理）
	UpdatedAt time.Time        // 最后一次更新时间（由GORM自动管理）
}

// 表名
func (c *Conversation) TableName() string {
	return "im_conversation"
}
