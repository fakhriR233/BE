package cartdto

type CreateCart struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	BookID int    `json:"book_id"`
	Qty       int    `json:"qty"`
	SubTotal  int    `json:"subtotal"`
	Status    string `json:"status"`
}

type UpdateCart struct {
	UserID    int `json:"user_id"`
	BookID int `json:"book_id"`
	Qty       int `json:"qty"`
	Subtotal  int `json:"subtotal"`
}

type CartRequest struct {
	UserID        int ` json:"user_id" `
	BookID     int ` json:"book_id" `
	Qty           int ` json:"qty" `
	Subtotal      int ` json:"subtotal"`
	TransactionID int ` json:"transaction_id"`
}

type CartResponse struct {
	ID       int `json:"id"`
	QTY      int `json:"qty"`
	SubTotal int `json:"subtotal"`
}