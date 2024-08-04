package entity

import "time"

type User struct {
	Id        int64     `gorm:"column:id;primary_key;auto_increment;" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255);default:'';" json:"name"`
	Password  string    `gorm:"column:password;type:varchar(255);default:'';" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:timestamp;default:NULL;" json:"deleted_at"`
}
