package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"colums:id; type:int; not null; primaryKey; autoIncrement; comment:'primary Key is ID'"`
	RoleName string `gorm:"colums:role_name"`
	IsActive bool   `gorm:"column:is_active; type:boolean;"`
	RoleNote string `gorm:"column:role_note; type:text;"`
}

func (u *Role) TableName() string {
	return "go_db_role"
}
