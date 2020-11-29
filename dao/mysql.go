package dao

import "github.com/jinzhu/gorm"

// shorturl
type Shorturl struct {
	gorm.Model
	Url string `json:"url"`
	Shorten	string `json:"shorten"`
}

func NewShorturl() Shorturl {
	return Shorturl{}
}

func (s *Shorturl) Create(req ShortenReq) (*ShortenResp, error) {
	return nil, nil
}