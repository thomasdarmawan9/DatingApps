package models

import (
	"final-project/helpers"
	"strconv"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type UserReq struct {
	Username   string `json:"username" example:"johndoe" valid:"required~Your username is required"`
	Email      string `json:"email" example:"user@example.com" valid:"required~Your email is required,email~Invalid email format"`
	FullName   string `json:"fullname" example:"John Doe" valid:"required~Your full name is required"`
	Address    string `json:"address" example:"123 Main St" valid:"required~Your address is required"`
	Password   string `json:"password" example:"password123" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age        int    `json:"age" example:"25" valid:"required~Your age is required,isUnderEight~Minimum of Age is 8 years old"`
	StatusUser string `json:"statusUser" example:"free" valid:"required~Your status user is required,statusUser~Status must be 'free' or 'premium'"`
}

type PhotoReq struct {
	Title    string `json:"title" example:"Beautiful Sunset" valid:"required~Title is required"`
	Caption  string `json:"caption" example:"A photo of a beautiful sunset"`
	PhotoUrl string `json:"photo_url" example:"https://example.com/photo.jpg" valid:"required~Url of your photo is required"`
	UserID   uint   `json:"user_id" example:"1"` // Optional: Direct relation to User
}

type RegisterReq struct {
	User  UserReq  `json:"user"`
	Photo PhotoReq `json:"photo"`
}

type LoginReq struct {
	Email    string `json:"email" example:"user@example.com" valid:"required,email"`
	Password string `json:"password" example:"password123" valid:"required"`
}

type SwipeProfileReq struct {
	ProfileID      uint `json:"profile_id" example:"1" binding:"required"`       // ID dari profile yang sedang login
	OtherProfileID uint `json:"other_profile_id" example:"2" binding:"required"` // ID dari profile yang di-swipe
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

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
