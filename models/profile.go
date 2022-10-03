package models

import "time"

type Profile struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment"`
	Phone  string `json:"phone" gorm:"type: varchar(255)"`
	Gender string `json:"gender" gorm:"type: varchar(255)"`
	Avatar string `json:"avatar" gorm:"type: text"`
	// UserID int    `json:"user_id"`
	// User      UserResponse 			`json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// TransactionID	[]int			`json:"transaction_id"`
	// Transaction   []TransactionResponse		`json:"purchasedBook" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ProfileResponse struct {
	Phone  string `json:"phone"`
	Gender string `json:"gender"`
	Avatar string `json:"address"`
	// UserID  int    `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
