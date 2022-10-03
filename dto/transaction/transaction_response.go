package transactiondto

import "_waysbook/models"

type TransactionUpdateResponse struct {
	ID							int							`json:"id"`
	UserID						int							`json:"user_id"`
	User 						models.UserResponse			`json:"user"`
	Attachment					string						`json:"attachment" form:"attachment" gorm:"type: varchar(255)"`
	BookID 						int                			`json:"book_id" form:"book_id" gorm:"-"`
	BooksPurchased   			models.BookResponse    		`json:"booksPurchased"`
	Total						int							`json:"totalPayment"`
	Status						string						`json:"status" form:"status" gorm:"type: varchar(255)"`
}
type TransactionDeleteResponse struct {
	ID							int							`json:"id"`
}