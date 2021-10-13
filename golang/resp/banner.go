package resp

type Banner struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	BannerID    string `json:"bannerId"`
	Url         string `json:"url"`
	RedirectUrl string `json:"redirectUrl"`
	OrderBy     int    `json:"orderBy"`
	IsDeleted   bool   `json:"isDeleted"`
	CreateAt    string `json:"createAt"`
	UpdateAt    string `json:"updateAt"`
}
