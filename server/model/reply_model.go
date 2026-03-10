package model

import "gorm.io/gorm"

type Reply struct {
	Base
	Content   string   `gorm:"not null" json:"content"`
	AccountId int      `json:"accountId"`
	ThreadId  int      `json:"threadId"`
	Account   *Account `gorm:"foreignKey:AccountId" json:"account"`
	Thread    *Thread  `gorm:"foreignKey:ThreadId" json:"thread"`
}

func (Reply) TableName() string {
	return "reply"
}

func (reply *Reply) AfterCreate(tx *gorm.DB) error {
	result := tx.
		Model(&Thread{}).
		Where("id = ?", reply.ThreadId).
		UpdateColumn("reply_count", gorm.Expr("reply_count + ?", 1))
	return result.Error
}

func (reply *Reply) AfterDelete(tx *gorm.DB) error {
	result := tx.
		Model(&Thread{}).
		Where("id = ?", reply.ThreadId).
		UpdateColumn("reply_count", gorm.Expr("reply_count - ?", 1))
	return result.Error
}

type ReplyPost struct {
	AccountId int    `json:"accountId" validate:"required"`
	ThreadId  int    `json:"threadId" validate:"required"`
	Content   string `json:"content" validate:"required"`
}

type ReplyPatch struct {
	Content string `json:"content"`
}

type ReplyInsert struct {
	AccountId int
	ThreadId  int
	Content   string
}

type ReplyUpdate struct {
	Id      int
	Content string
}
