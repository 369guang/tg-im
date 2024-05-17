package model

import "time"

type Group struct {
	ID           uint      `gorm:"primarykey"`   // ID
	Name         string    `json:"name"`         // 群组名称
	Avatar       string    `json:"avatar"`       // 群组头像
	Announcement string    `json:"announcement"` // 群公告
	OwnerID      uint      `json:"owner_id"`     // 群主ID
	CreatedAt    time.Time `json:"created_at"`   // 创建时间
	UpdatedAt    time.Time `json:"updated_at"`   // 更新时间
}

// 表名
func (g *Group) TableName() string {
	return "im_group"
}

type GroupMember struct {
	ID        uint      `gorm:"primarykey"`            // ID
	GroupID   uint      `gorm:"index" json:"group_id"` // 群组ID
	UserID    uint      `gorm:"index" json:"user_id"`  // 用户ID
	Nickname  string    `json:"nickname"`              // 用户在群里的昵称
	Role      int       `gorm:"index" json:"role"`     // 角色，1：普通成员，2：管理员
	CreatedAt time.Time `json:"created_at"`            // 创建时间
}

// 表名
func (gm *GroupMember) TableName() string {
	return "im_group_member"
}
