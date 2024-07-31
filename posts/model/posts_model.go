package model

type Posts struct {
	Id    int    `json:"id" gorm:"primary_key"`
	Title string `json:"name" gorm:"type:varchar(100);not null"`
}
