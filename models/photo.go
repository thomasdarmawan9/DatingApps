package models

type Photo struct {
	GormModel
	Title    string `json:"title" form:"title" valid:"required~Title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url" valid:"required~Url of your photo is required"`
	UserID   uint   `json:"user_id" form:"user_id"`
	User     User   `json:"user" form:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
