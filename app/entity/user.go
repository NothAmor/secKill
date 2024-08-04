package entity

import (
	"time"
)

type User struct {
	Id        int64     `gorm:"column:id;primary_key;auto_increment;" json:"id"` // 用户ID
	UUID      string    `json:"uuid" gorm:"index;comment:用户UUID"`                // 用户UUID
	Username  string    `json:"userName" gorm:"index;comment:用户登录名"`             // 用户登录名
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`                        // 用户登录密码
	NickName  string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`       // 用户昵称
	HeaderImg string    `json:"headerImg" gorm:"default:'';comment:用户头像"`        // 用户头像
	Phone     string    `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email     string    `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable    int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:timestamp;default:NULL;" json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}
