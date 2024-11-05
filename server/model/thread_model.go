package model

type Thread struct {
	Base
	Title      string    `gorm:"unique;not null" json:"title"`
	Content    string    `gorm:"not null" json:"content"`
	AccountID  int       `json:"accountId"`
	SubforumID int       `json:"subforumId"`
	Account    *Account  `gorm:"foreignKey:AccountID" json:"account"`
	Subforum   *Subforum `gorm:"foreignKey:SubforumID" json:"subforum"`
	Replies    []*Reply  `gorm:"foreignKey:ThreadID" json:"replies"`
}

func (Thread) TableName() string {
	return "thread"
}

type ThreadPost struct {
	AccountID  int    `json:"accountId" binding:"required"`
	SubforumID int    `json:"subforumId" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

type ThreadPatch struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ThreadInsert struct {
	AccountID  int
	SubforumID int
	Title      string
	Content    string
}

type ThreadUpdate struct {
	ID      int
	Title   string
	Content string
}
