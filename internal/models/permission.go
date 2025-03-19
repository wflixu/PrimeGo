package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	PermissionCode string `gorm:"uniqueIndex;not null;size:100"`
	Description    string
	Roles          []Role `gorm:"many2many:role_permissions;"`
}
