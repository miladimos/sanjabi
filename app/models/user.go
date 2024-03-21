package models

// type Role struct {
// 	ID   int    `gorm:"column:id; primary_key; not null" json:"id"`
// 	Role string `gorm:"column:role" json:"role"`
// 	BaseModel
// }

type User struct {
	ID        int64  `gorm:"column:id; primary_key:auto_increment; not null" json:"id"`
	Firstname string `gorm:"column:firstname" json:"firstname"`
	Lastname  string `gorm:"column:lastname" json:"lastname"`
	Email     string `gorm:"column:email;unique" json:"email"`
	Password  string `gorm:"column:password" json:"-"`
	// BaseModel
}
