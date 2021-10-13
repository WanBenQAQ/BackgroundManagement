package model

type Banner struct {
	BannerID    string `json:"bannerID" gorm:"column:banner_id"`
	Url         string `json:"url" gorm:"column:url"`
	RedirectUrl string `json:"redirectUrl" gorm:"column:redirect_url"`
	OrderBy     int    `json:"orderBy" gorm:"column:order_by"`
	CreateUser  string `json:"createUser" gorm:"column:create_user"`
	UpdateUser  string `json:"updateUser" gorm:"column:update_user"`
	IsDeleted   bool   `json:"isDeleted" gorm:"column:is_deleted"`
	CreateAt    string `json:"createAt" gorm:"column:create_at"`
	UpdateAt    string `json:"updateAt" gorm:"column:update_at"`
}
