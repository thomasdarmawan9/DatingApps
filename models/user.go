package models

import (
	"final-project/helpers"
	"strconv"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for an user
type User struct {
	GormModel
	Username   string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email      string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	FullName   string `gorm:"not null" json:"fullname" form:"fullname" valid:"required~Your full name is required"`
	Address    string `gorm:"not null" json:"address" form:"address" valid:"required~Your address is required"`
	Password   string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age        uint   `gorm:"not null;" json:"age" form:"age" valid:"required~Your age is required,isUnderEight~Minimum of Age is 8 years old"`
	StatusUser string `gorm:"not null" json:"status_user" form:"status_user" valid:"required~Your status user is required,statusUser~Status must be 'free' or 'premium'"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Add custom validator for isUnderEight
	govalidator.TagMap["isUnderEight"] = govalidator.Validator(func(str string) bool {
		num, _ := strconv.Atoi(str)
		return num > 7
	})

	// Add custom validator for statusUser
	govalidator.TagMap["statusUser"] = govalidator.Validator(func(str string) bool {
		return str == "free" || str == "premium"
	})

	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
