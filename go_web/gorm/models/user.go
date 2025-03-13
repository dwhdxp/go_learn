package models

import (
	"time"
)

// 表名默认为结构体复数users
type User struct {
	ID        uint      `gorm:"auto_increment;primary key" json:"-"` // 默认主键
	Name      string    `gorm:"size:64;not null" json:"name"`
	Age       uint8     `gorm:"default:18"`
	Email     string    `gorm:"unique" json:"email"`
	CreatedAt time.Time `gorm:"column:create_time" json:"-"` // 使用tag中的column来指定表名
	UpdatedAt time.Time `gorm:"column:update_time" json:"-"`
}

// 将User表名设置为`user`
func (User) TableName() string { return "user" }
