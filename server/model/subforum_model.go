package model

type Subforum struct {
	Base
	Title       string    `gorm:"unique;not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Threads     []*Thread `gorm:"foreignKey:SubforumID" json:"threads"`
}

func (Subforum) TableName() string {
	return "subforum"
}

type SubforumPost struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
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
	ID          int
	Title       string
	Description string
}
