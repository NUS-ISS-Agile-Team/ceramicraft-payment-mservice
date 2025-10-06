package model

import "time"

type UserAccount struct {
	ID        int       `gorm:"primaryKey"`
	UserId    int       `gorm:"uniqueIndex;not null"`
	AccountNo string    `gorm:"uniqueIndex;not null"`
	Balance   int       `gorm:"not null;default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (u *UserAccount) TableName() string {
	return "user_accounts"
}

func (u *UserAccount) GetHiddenAccountNo() string {
	if len(u.AccountNo) <= 8 {
		return u.AccountNo
	}
	return u.AccountNo[0:4] + "****" + u.AccountNo[len(u.AccountNo)-4:]
}
