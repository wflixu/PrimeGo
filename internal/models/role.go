package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName    string `gorm:"uniqueIndex;not null"`
	Description string
	Users       []User       `gorm:"many2many:user_roles;"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`
}
