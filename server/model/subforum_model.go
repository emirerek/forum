package model

type Subforum struct {
	Base
	Title       string    `gorm:"unique;not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Threads     []*Thread `gorm:"foreignKey:SubforumId" json:"threads"`
	ThreadCount int       `gorm:"default:0" json:"threadCount"`
}

func (Subforum) TableName() string {
	return "subforum"
}

type SubforumPost struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type SubforumPatch struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type SubforumInsert struct {
	Title       string
	Description string
}

type SubforumUpdate struct {
	Id          int
	Title       string
	Description string
}
