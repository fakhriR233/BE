package models

import "time"

type Book struct {
	ID              int    `json:"id" gorm:"primary_key:auto_increment"`
	Title           string `json:"title" form:"title" gorm:"type: varchar(255)"`
	PublicationDate string `json:"publicationDate" form:"publication_date" gorm:"type: varchar(255)"`
	Pages           int    `json:"pages" form:"pages" gorm:"type: int"`
	ISBN            string `json:"ISBN" form:"ISBN" gorm:"type: varchar(255)"`
	Author          string `json:"author" form:"author" gorm:"type: varchar(255)"`
	Price           int    `json:"price" form:"price" gorm:"type: int"`
	Description     string `json:"description" gorm:"type:text" form:"description"`
	BookAttachment  string `json:"bookAttachment" form:"Attachment" gorm:"type: varchar(255)"`
	Thumbnail       string `json:"thumbnail" form:"Thumbnail" gorm:"type: varchar(255)"`
	// UserID          int          `json:"user_id"`
	// User            UserResponse `json:"user"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type BookResponse struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	PublicationDate string `json:"publicationDate"`
	Pages           int    `json:"pages"`
	ISBN            int    `json:"ISBN"`
	Author          string `json:"author"`
	Price           int    `json:"price"`
	Description     string `json:"description" gorm:"type: text"`
	BookAttachment  string `json:"bookAttachment"`
	Thumbnail       string `json:"thumbnail"`
	// UserID          int          `json:"-"`
	// User            UserResponse `json:"user"`
}

type BookTransactionResponse struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	PublicationDate string `json:"publicationDate"`
	Pages           int    `json:"pages"`
	ISBN            int    `json:"ISBN"`
	Author          string `json:"author"`
	Price           int    `json:"price"`
	Description     string `json:"description" gorm:"type: text"`
	BookAttachment  string `json:"bookAttachment"`
	Thumbnail       string `json:"thumbnail"`
	// UserID          int          `json:"-"`
	// User            UserResponse `json:"user"`
}

// type BookUserResponse struct {
// 	ID              int          `json:"id"`
// 	Title           string       `json:"title"`
// 	PublicationDate string       `json:"publicationDate"`
// 	Pages           int          `json:"pages"`
// 	ISBN            int          `json:"ISBN"`
// 	Author          string       `json:"author"`
// 	Price           int          `json:"price"`
// 	Description     string       `json:"description" gorm:"type: text"`
// 	BookAttachment  string       `json:"bookAttachment"`
// 	Thumbnail       string       `json:"thumbnail"`
// 	UserID          int          `json:"user_id"`
// 	User            UserResponse `json:"user"`
// 	// UserID              int               `json:"-"`
// 	CreatedAt time.Time `json:"-"`
// 	UpdatedAt time.Time `json:"-"`
// }

func (BookResponse) TableName() string {
	return "books"
}

func (BookTransactionResponse) TableName() string {
	return "books"
}
