package usersdto

import "_waysbook/models"

type UserResponse struct {
	ID						int		`json:"id"`
	FullName     			string `json:"FullName" form:"name" `
	Email    				string `json:"email" form:"email" `
	Password 				string `json:"password" form:"password" `
	Gender 					string `json:"gender" form:"gender" `
	Phone 					int `json:"phone" form:"phone" `
	Avatar 					string `json:"avatar" form:"avatar" `
	BooksPurchased 			models.BookResponse `json:"booksPurchased"`
  }

type UserResponseDelete struct {
	ID				int		`json:"id"`
  }