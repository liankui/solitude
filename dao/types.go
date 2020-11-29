package dao

type ShortenReq struct {
	Url string `form:"url"`
}

type ShortenResp struct {
	Shorten string `json:"shorten"`
}
