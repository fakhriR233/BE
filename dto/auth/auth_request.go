package authdto

type AuthRequest struct {
  ID				  	int			`json:"id" gorm:"primary_key:auto_increment"`
  FullName     			string `json:"fullname" form:"fullname" validate:"required"`
  Email    				string `json:"email" form:"email" validate:"required"`
  Password 				string `json:"password" form:"password" validate:"required"`
  Gender 				string `json:"gender" form:"gender" validate:"required"`
  Phone 				string `json:"phone" form:"phone" validate:"required"`
  Address 				string `json:"address" form:"address" validate:"required"`
}

type RegisterRequest struct {
  FullName     	string `json:"fullname" form:"fullname" validate:"required"`
  Email    		string `json:"email" form:"email" validate:"required"`
  Password 		string `json:"password" form:"password" validate:"required"`
}

type LoginRequest struct {
  Email    		string `json:"email" form:"email" validate:"required"`
  Password 		string `json:"password" form:"password" validate:"required"`
}