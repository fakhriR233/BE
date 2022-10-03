package models

import "time"

type Transaction struct {
	ID         int          `json:"id" gorm:"primary_key:auto_increment"`
	UserID     int          `json:"user_id"`
	User       UserResponse `json:"user"`
	Attachment string       `json:"attachment" form:"attachment" gorm:"type: varchar(255)"`
	Cart	[]Cart			`json:"cart"`
	// BookID		int			`json:"book_id"`
	// Book		BookResponse		`json:"books" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Total      int          `json:"totalPayment"`
	Status     string       `json:"status" gorm:"type:varchar(255)"`
	CreatedAt  time.Time    `json:"-"`
	UpdatedAt  time.Time    `json:"-"`
}

type TransactionResponse struct {
	ID     int          `json:"-"`
	UserID int          `json:"-"`
	User   UserResponse `json:"user"`
	// User      				ProfileResponse    		`json:"-" gorm:"foreignKey:UserID"`
	BookID    int       		`json:"-"`
	Book      BookResponse      `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
