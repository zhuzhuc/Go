package model

type Blog struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"Title" gorm:"not null;column:title;size:225"`
	Post     string `json:"Post" gorm:"not null;column:post;size:225"`
	Image    string `json:"Image" gorm:"column:image;size:255"`
	AuthorID uint   `json:"AuthorID" gorm:"column:author_id"`
	Author   string `json:"Author" gorm:"column:author;size:100"`
}
