package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" valid:"required~Url of your social media is required"`
	UserID         uint
	User           *User
	ProfileID      uint    `json:"profile_id" form:"profile_id"`                                                 // Foreign key to Profile
	Profile        Profile `json:"profile" form:"profile" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relasi ke Profile
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
