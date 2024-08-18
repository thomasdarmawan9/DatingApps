package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Profile struct {
	GormModel
	UserID       uint           `json:"user_id" form:"user_id"`
	User         User           `json:"user" form:"user" valid:"required" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedias []SocialMedia  `json:"social_medias" form:"social_medias" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalSwipe   uint           `json:"total_swipe" form:"total_swipe" gorm:"default:0"`
	Matches      []MatchProfile `json:"matches" form:"matches" gorm:"foreignKey:ProfileID"`
}

func (p *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

type MatchProfile struct {
	GormModel
	ProfileID      uint `json:"profile_id" form:"profile_id"`             // Profile yang sedang login
	OtherProfileID uint `json:"other_profile_id" form:"other_profile_id"` // Profile lain yang di-swipe
	StatusMatch    int  `json:"status_match" form:"status_match"`         // Status match (0: not match, 1: match)"
}

func (mp *MatchProfile) BeforeCreate(tx *gorm.DB) (err error) {
	if mp.StatusMatch != 0 && mp.StatusMatch != 1 {
		err = gorm.ErrInvalidData
		return
	}

	err = nil
	return
}
