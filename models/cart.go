package models

import "time"

type Cart struct {
	ID        int          `json:"id" gorm:"primary_key:auto_increment"`
	QTY       int          `json:"qty"`
	SubTotal  int          `json:"subtotal"`
	BookID    int          `json:"book_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Book      BookResponse `json:"book" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TransactionID    int          `json:"transaction_id"`
	Transaction      TransactionResponse `json:"transaction"`
	Status    string       `json:"status"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
}

type CartResponse struct {
	ID       int          `json:"id"`
	Qty      int          `json:"qty"`
	SubTotal int          `json:"subtotal"`
	BookID   int          `json:"book_id"`
	Book     BookResponse `json:"book"`
	Status   string       `json:"status"`
}

func (CartResponse) TableName() string {
	return "carts"
}
