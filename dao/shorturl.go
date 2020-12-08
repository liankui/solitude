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

func (s *Shorturl) GetShorten(url, shorten string) (string, error) {
	// 先从redis里查，如果没有则从mysql中查。如果都没有则存mysql和redis
	hGet, _ := Redis.HGet("longurl", url).Result()
	if hGet != "" {
		return hGet, nil
	}

	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return "", err
	}

	DB.Where("url = ?", url).Last(&s)

	if s.ID == 0 {
		s.Url = url
		s.Shorten = shorten
		if err := tx.Table(s.TableName()).Create(&s).Error; err != nil {
			tx.Rollback()
			return "", err
		}
	}

	if err := Redis.HSet("shorturl", url, s.Shorten).Err(); err != nil {
		tx.Rollback()
		return "", err
	}
	if err := Redis.HSet("longurl", s.Shorten, url).Err(); err != nil {
		tx.Rollback()
		return "", err
	}

	return s.Shorten, tx.Commit().Error
}

func (s *Shorturl) GetUrl(shorten string) (string, error) {
	hGet, _ := Redis.HGet("longurl", shorten).Result()
	if hGet == "" {
		err := DB.Where("shorten = ?", shorten).Last(&s).Error
		if err != nil {
			return "", err
		}
		return s.Url, nil
	} else {
		return hGet, nil
	}
}
