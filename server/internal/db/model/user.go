package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName  string   `gorm:"type:varchar(20);not null;unique" json:"user_name"`                // 用户名
	Password  string   `gorm:"type:varchar(20);not null" json:"password"`                        // 密码
	Email     string   `gorm:"type:varchar(50);unique" json:"email"`                             // 邮箱
	Phone     string   `gorm:"type:varchar(20);unique" json:"phone"`                             // 手机号码
	Sex       int      `gorm:"type:int;default:1" json:"sex"`                                    // 性别，1：男，2：女
	Avatar    string   `gorm:"type:varchar(255)" json:"avatar"`                                  // 头像
	Status    int      `gorm:"type:int;default:1" json:"status"`                                 // 状态，1：正常，2：禁用
	Role      int      `gorm:"type:int;default:1" json:"role"`                                   // 角色，1：普通用户，2：管理员
	LastLogin int      `gorm:"type:int" json:"last_login"`                                       // 最后登录时间
	UserInfo  UserInfo `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 用户信息
}

// 表名
func (u *User) TableName() string {
	return "im_user"
}

type UserInfo struct {
	ID        uint   `gorm:"primarykey"`                         // ID
	UserID    uint   `gorm:"index" json:"user_id"`               // 用户ID
	RealName  string `gorm:"type:varchar(20)" json:"real_name"`  // 真实姓名
	NickName  string `gorm:"type:varchar(20)" json:"nick_name"`  // 昵称
	Birthday  string `gorm:"type:varchar(20)" json:"birthday"`   // 生日
	Address   string `gorm:"type:varchar(255)" json:"address"`   // 地址
	Signature string `gorm:"type:varchar(255)" json:"signature"` // 个性签名
}

// 表名
func (u *UserInfo) TableName() string {
	return "im_user_info"
}
