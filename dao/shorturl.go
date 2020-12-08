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

func (s *Shorturl) Insert(url, shorten string) (string, error) {

	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return "", err
	}

	tx.Where("url = ?", url).Last(&s)

	if s.ID == 0 {
		s.Url = url
		s.Shorten = shorten
		if err := tx.Table(s.TableName()).Create(&s).Error; err != nil {
			tx.Rollback()
			return "", err
		}
		if err := Redis.HSet("shorturl", shorten, url).Err(); err != nil {
			tx.Rollback()
			return "", err
		}
	}

	return s.Shorten, tx.Commit().Error
}

