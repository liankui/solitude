package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/liankui/solitude/dao"
	"math/rand"
	"time"
)

func Shorten(c *gin.Context) {
	url := c.Query("url")
	shorten := RandStringRunes(6)

	// 持久化到mysql，存入redis中，返回给调用者短链
	s := dao.NewShorturl()
	str, err := s.GetShorten(url, shorten)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "insert mysql error" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": str,
	})
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Expand(c *gin.Context) {
	shorten := c.Param("shorten")
	// 使用者访问短链，查短链映射长链，并重定向至长链
	s := dao.NewShorturl()
	getUrl, err := s.GetUrl(shorten)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "get url error" + err.Error(),
		})
		return
	}
	c.String(200, getUrl)
	//c.Redirect(http.StatusMovedPermanently, getUrl)
}


