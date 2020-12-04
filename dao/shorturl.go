package dao

import (
	"github.com/jinzhu/gorm"
	"time"
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

func (s *Shorturl) Insert(url, shorten string) (Shorturl, error) {
	s.Url = url
	s.Shorten = shorten
	err := DB.TestDate.Table(s.TableName()).Create(&s).Error
	return *s, err
}

type TestTable struct {
	ID int `json:"id"`
	StartDate time.Time `json:"start_date"`
}

func FindTestData() ([]TestTable, error) {
	var t []TestTable
	err := DB.TestDate.Table("test_table").Find(&t).Error
	return t, err
}
