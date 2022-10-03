package bookdto

type BookRequest struct {
	ID        					int       					`json:"id" gorm:"primary_key:auto_increment"`
  Title      					string						`json:"title" gorm:"type: varchar(255)"`
  PublicationDate      			string    					`json:"publicationDate" gorm:"type: varchar(255)"`
  Pages     					int    						`json:"pages" gorm:"type: int"`
  ISBN     						string    						`json:"ISBN" gorm:"type: varchar(255)"`
  Author						string						`json:"author" gorm:"type: varchar(255)"`
  Price     					int    						`json:"price" gorm:"type: int"`
  Description					string						`json:"description" gorm:"type:text" form:"desc"`
  BookAttachment				string						`json:"bookAttachment" gorm:"type: varchar(255)"`
  Thumbnail						string						`json:"thumbnail" gorm:"type: varchar(255)"`
}

type BookUpdateRequest struct {
	ID        					int       					`json:"id"`
  Title      					string						`json:"title" gorm:"type: varchar(255)"`
  PublicationDate      			string    					`json:"publicationDate" gorm:"type: varchar(255)"`
  Pages     					int    						`json:"pages" gorm:"type: int"`
  ISBN     						string    						`json:"ISBN" gorm:"type: varchar(255)"`
  Author						string						`json:"author" gorm:"type: varchar(255)"`
  Price     					int    						`json:"price" gorm:"type: int"`
  Description					string						`json:"description" gorm:"type:text" form:"desc"`
  BookAttachment				string						`json:"bookAttachment" gorm:"type: varchar(255)"`
  Thumbnail						string						`json:"thumbnail" gorm:"type: varchar(255)"`
}