package model

import "time"

type Friend struct {
	ID        uint      `gorm:"primarykey"`             // ID
	UserID    uint      `gorm:"index" json:"user_id"`   // 用户ID
	FriendID  uint      `gorm:"index" json:"friend_id"` // 好友ID
	CreatedAt time.Time `json:"created_at"`             // 创建时间
}

// 表名
func (f *Friend) TableName() string {
	return "im_friend"
}
