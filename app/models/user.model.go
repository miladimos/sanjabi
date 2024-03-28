package models

import "gorm.io/gorm"

// type Role struct {
// 	ID   int    `gorm:"column:id; primary_key; not null" json:"id"`
// 	Role string `gorm:"column:role" json:"role"`
// 	BaseModel
// }

type User struct {
	Firstname string `gorm:"column:firstname" json:"firstname"`
	Lastname  string `gorm:"column:lastname" json:"lastname"`
	Email     string `gorm:"column:email;unique" json:"email"`
	Mobile    string `gorm:"column:mobile;unique" json:"mobile"`
	Password  string `gorm:"column:password" json:"-"`
	gorm.Model
}
