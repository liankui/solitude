package logic

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/liankui/solitude/dao"
	"math/rand"
)

func Shorten(c *gin.Context) {
	url := c.Query("url")

	hello := "123456"
	shorten := base64.URLEncoding.EncodeToString([]byte(hello))

	s := dao.NewShorturl()
	insert, err := s.Insert(url, shorten)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "insert error" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": insert,
	})
}

func base64Encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	l := len(letterRunes)
	for i := range b {
		b[i] = letterRunes[rand.Intn(l)]
	}
	return string(b)
}

func Expand(c *gin.Context) {
	data, _ := dao.FindTestData()

	c.JSON(200, gin.H{
		"message": data,
	})
}


