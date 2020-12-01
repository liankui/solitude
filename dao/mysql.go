package dao

import "github.com/jinzhu/gorm"

// shorturl
type Shorturl struct {
	gorm.Model
	Url     string `json:"url"`
	Shorten string `json:"shorten"`
}

func NewShorturl() Shorturl {
	return Shorturl{}
}

func (s *Shorturl) Insert(url, shorten string) (Shorturl, error) {
	s.Url = url
	s.Shorten = shorten
	err := DB.TestDate.Create(&s).Error
	return *s, err
}
