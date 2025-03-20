package models

type Article struct{
	ID uint `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"unique;not null"`
	Article string `gorm:"not null"`
}

type ArticleSelect struct{
	Name string `json:"name"`
	Article string `json:"article"`
}

type ArticlePagination struct{
	Data []ArticleSelect
	TotalRows int64
	TotalPage int
	Page int
	Limit int
}