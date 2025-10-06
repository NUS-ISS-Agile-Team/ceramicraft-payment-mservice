package model

import "time"

type RedeemCode struct {
	ID         uint      `gorm:"primaryKey"`
	Code       string    `gorm:"uniqueIndex;not null"`
	Amount     int       `gorm:"not null"`
	UsedUserId int       `gorm:"not null;default:0"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type RedeemCodeQuery struct {
	Code  *string
	Used  *bool
	Limit int
}
