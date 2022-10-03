package transactiondto

import "_waysbook/models"

type TransactionRequest struct {
	ID							int							`json:"id" gorm:"primary_key:auto_increment"`
	UserID						int							`json:"user_id"`
	User 						models.UserResponse			`json:"user"`
	Attachment					string						`json:"attachment" form:"attachment" gorm:"type: varchar(255)"`
	BookID 						int                			`json:"book_id" form:"book_id" gorm:"-"`
	Books   					models.BookResponse    		`json:"booksPurchased"`
	Total						int							`json:"totalPayment"`
	Status						string						`json:"status" form:"status" gorm:"type: varchar(255)"`
}

type TransactionUpdateRequest struct {
	UserID						int							`json:"user_id"`
	User 						models.UserResponse			`json:"user"`
	Attachment					string						`json:"attachment" form:"attachment" gorm:"type: varchar(255)"`
	BookID 						int                			`json:"book_id" form:"book_id" gorm:"-"`
	Books   					models.BookResponse    		`json:"booksPurchased"`
	Total						int							`json:"totalPayment"`
	Status						string						`json:"status" form:"status" gorm:"type: varchar(255)"`
}