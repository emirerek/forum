package model

type Reply struct {
	Base
	Content   string   `gorm:"not null" json:"content"`
	AccountID int      `json:"accountId"`
	ThreadID  int      `json:"threadId"`
	Account   *Account `gorm:"foreignKey:AccountID" json:"account"`
	Thread    *Thread  `gorm:"foreignKey:ThreadID" json:"thread"`
}

func (Reply) TableName() string {
	return "reply"
}

type ReplyPost struct {
	AccountID int    `json:"accountId" binding:"required"`
	ThreadID  int    `json:"threadId" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

type ReplyPatch struct {
	Content string `json:"content"`
}

type ReplyInsert struct {
	AccountID int
	ThreadID  int
	Content   string
}

type ReplyUpdate struct {
	ID      int
	Content string
}
