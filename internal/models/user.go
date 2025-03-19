package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	RealName string
	Email    string `gorm:"uniqueIndex"`
	IsActive bool   `gorm:"default:true"`
	Roles    []Role `gorm:"many2many:user_roles;"`
}

type UserRole struct {
	UserID uint `gorm:"primaryKey"`
	RoleID uint `gorm:"primaryKey"`
}
