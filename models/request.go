package models

type UserReq struct {
	Username   string `json:"username" example:"johndoe"`
	Email      string `json:"email" example:"user@example.com"`
	FullName   string `json:"fullname" example:"John Doe"`
	Address    string `json:"address" example:"123 Main St"`
	Password   string `json:"password" example:"password123"`
	Age        int    `json:"age" example:"25"`
	StatusUser string `json:"statusUser" example:"free"`
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
