package model

import "time"

type Message struct {
	ID             uint      `gorm:"primarykey"`                   // ID
	ConversationID uint      `gorm:"index" json:"conversation_id"` // 会话ID
	SenderID       uint      `json:"sender_id"`                    // 发送者ID
	MessageType    int       `json:"message_type"`                 // 消息类型，1：文本，2：图片，3：文件, 4: 语音, 5: 视频
	Content        string    `json:"content"`                      // 消息内容
	Status         int       `json:"status"`                       // 消息状态，1：已发送，2：已送达，3：已读, 4: 撤回
	CreatedAt      time.Time // 创建时间（由GORM自动管理）
	UpdatedAt      time.Time // 最后一次更新时间（由GORM自动管理）
}

// 表名
func (m *Message) TableName() string {
	return "im_message"
}
