package dao

import (
	"github.com/jinzhu/gorm"
)

// shorturl
type Shorturl struct {
	gorm.Model
	Url     string `json:"url"`
	Shorten string `json:"shorten"`
}

func (Shorturl) TableName() string {
	return "shorturl"
}

func NewShorturl() Shorturl {
	return Shorturl{}
}

func (s *Shorturl) Insert(url, shorten string) error {
	s.Url = url
	s.Shorten = shorten
	err := DB.TestDate.Table(s.TableName()).Create(&s).Error
	return err
}

// redis----------
func (s *Shorturl) InsertRedis(url, shorten string) error {
	err := Redis.HSet("shorturl", shorten, url).Err()
	return err
}
