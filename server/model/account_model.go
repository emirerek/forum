package model

type Account struct {
	Base
	Username     string    `gorm:"unique;not null" json:"username"`
	Email        string    `gorm:"unique;not null" json:"email"`
	ProfilePath  string    `gorm:"" json:"profilePath"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Threads      []*Thread `gorm:"foreignKey:AccountId" json:"threads"`
	Replies      []*Reply  `gorm:"foreignKey:AccountId" json:"replies"`
	IsAdmin      bool      `gorm:"default:false" json:"isAdmin"`
}

func (Account) TableName() string {
	return "account"
}

type AccountCredentials struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	IsAdmin      bool   `json:"isAdmin"`
}

func (AccountCredentials) TableName() string {
	return "account"
}

type AccountPost struct {
	Username    string `json:"username" form:"username" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	ProfilePath string `json:"profilePath" form:"profilePath"`
	Password    string `json:"password" form:"password" validate:"required"`
}

type AccountPostLogin struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type AccountPatchCredentials struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type AccountPatchPassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type AccountInsert struct {
	Username     string
	Email        string
	ProfilePath  string
	PasswordHash string
	IsAdmin      bool // default false
}

type AccountUpdateCredentials struct {
	Id       int
	Username string
	Email    string
}

type AccountUpdatePassword struct {
	PasswordHash string
}
