package models

// User represents the model for an user
type User struct {
	GormModel
	Username   string  `gorm:"not null;uniqueIndex" json:"username" form:"username"`
	Email      string  `gorm:"not null;uniqueIndex" json:"email" form:"email"`
	FullName   string  `gorm:"not null" json:"fullname" form:"fullname"`
	Address    string  `gorm:"not null" json:"address" form:"address"`
	Password   string  `gorm:"not null" json:"password" form:"password"`
	Age        uint    `gorm:"not null;" json:"age" form:"age"`
	StatusUser string  `gorm:"not null" json:"status_user" form:"status_user"`
	Photos     []Photo `json:"photos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relasi one-to-many dengan Photo
}
