package model

type Account struct {
	Base
	Username     string    `gorm:"unique;not null" json:"username"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"-" json:"-"`
	Threads      []*Thread `gorm:"foreignKey:AccountID" json:"threads"`
	Replies      []*Reply  `gorm:"foreignKey:AccountID" json:"replies"`
}

func (Account) TableName() string {
	return "account"
}

type AccountCredentials struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

type AccountPost struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AccountPatchCredentials struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type AccountPatchPassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type AccountInsert struct {
	Username     string
	Email        string
	PasswordHash string
}

type AccountUpdateCredentials struct {
	ID       int
	Username string
	Email    string
}

type AccountUpdatePassword struct {
	PasswordHash string
}
