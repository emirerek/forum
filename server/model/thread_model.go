package model

import "gorm.io/gorm"

type Thread struct {
	Base
	Title      string    `gorm:"unique;not null" json:"title"`
	Content    string    `gorm:"not null" json:"content"`
	AccountId  int       `json:"accountId"`
	SubforumId int       `json:"subforumId"`
	Account    *Account  `gorm:"foreignKey:AccountId" json:"account"`
	Subforum   *Subforum `gorm:"foreignKey:SubforumId" json:"subforum"`
	Replies    []*Reply  `gorm:"foreignKey:ThreadId" json:"replies"`
	ReplyCount int       `gorm:"default:0" json:"replyCount"`
}

func (Thread) TableName() string {
	return "thread"
}

func (thread *Thread) AfterCreate(tx *gorm.DB) error {
	result := tx.
		Model(&Subforum{}).
		Where("id = ?", thread.SubforumId).
		UpdateColumn("thread_count", gorm.Expr("thread_count + ?", 1))
	return result.Error
}

func (thread *Thread) AfterDelete(tx *gorm.DB) error {
	result := tx.
		Model(&Subforum{}).
		Where("id = ?", thread.SubforumId).
		UpdateColumn("thread_count", gorm.Expr("thread_count - ?", 1))
	return result.Error
}

type ThreadPost struct {
	AccountId  int    `json:"accountId" validate:"required"`
	SubforumId int    `json:"subforumId" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
}

type ThreadPatch struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ThreadInsert struct {
	AccountId  int
	SubforumId int
	Title      string
	Content    string
}

type ThreadUpdate struct {
	Id      int
	Title   string
	Content string
}
