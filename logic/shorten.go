package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/liankui/solitude/dao"
	"math/rand"
	"net/http"
)

func Shorten(c *gin.Context) {
	url := c.Query("url")
	shorten := RandStringRunes(5)

	// 持久化到mysql，存入redis中，返回给调用者短链
	s := dao.NewShorturl()
	err := s.Insert(url, shorten)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "insert error" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": shorten,
	})
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	l := len(letterRunes)
	for i := range b {
		b[i] = letterRunes[rand.Intn(l)]
	}
	return string(b)
}

func Expand(c *gin.Context) {
	url := c.Param("url")
	// 使用者访问短链，查短链映射长链，并重定向至长链

	c.Redirect(http.StatusMovedPermanently, url)
	c.JSON(200, gin.H{
		"message": url,
	})
}


