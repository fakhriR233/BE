package authdto

type AuthResponse struct {
  FullName     string `gorm:"type: varchar(255)" json:"FullName"`
  Email       string `gorm:"type: varchar(255)" json:"email"`
  Password    string `gorm:"type: varchar(255)" json:"password"`
  Token    	string `gorm:"type: varchar(255)" json:"token"`
}

type LoginResponse struct {
  FullName     string `gorm:"type: varchar(255)" json:"FullName"`
  Email       string `gorm:"type: varchar(255)" json:"email"`
  Status       string `gorm:"type: varchar(255)" json:"status"`
  Token    	string `gorm:"type: varchar(255)" json:"token"`
}

type RegisterResponse struct {
  Email       string `gorm:"type: varchar(255)" json:"email"`
  Token    	string `gorm:"type: varchar(255)" json:"token"`
}

type CheckAuthResponse struct {
  ID				  int			`json:"id"`
  FullName     string `gorm:"type: varchar(255)" json:"FullName"`
  Email       string `gorm:"type: varchar(255)" json:"email"`
  Token    	string `gorm:"type: varchar(255)" json:"token"`
  Status    	string `gorm:"type: varchar(255)" json:"status"`
  Gender 	  	string		`json:"gender" gorm:"type: varchar(255)"`
  Phone 	 	string		`json:"phone" gorm:"type: varchar(255)"`
  Avatar 	  	string		`json:"address" gorm:"type: varchar(255)"`
  // Subscribe 	string		`json:"subscribe" gorm:"type: varchar(255)"`
}