package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/liankui/solitude/dao"
	"math/rand"
	"time"
)

// Shorten godoc
// @Summary 返回短链
// @Description 使用url参数传递长链接，接口返回短链接
// @Tags url
// @Accept application/json
// @Produce application/json
// @Param url query string false "string valid"
// @Success 200 {string} string "短链接字符串"
// @Failure 500 {string} string "insert mysql error"
// @Router /shorten [get]
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

// Expand godoc
// @Summary 使用短链接
// @Description 使用短链接跳转至原链接
// @Tags url
// @Accept application/json
// @Produce application/json
// @Param shorten path string true "短链接"
// @Success 200 {string} string "短链接字符串"
// @Failure 500 {string} string "get url error"
// @Router /expand/{shorten} [get]
func Expand(c *gin.Context) {
	shorten := c.Param("shorten")
	// 使用者访问短链，查短链映射长链，并重定向至长链
	s := dao.NewShorturl()
	getUrl, err := s.GetUrl(shorten)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "get url error: " + err.Error(),
		})
		return
	}
	c.String(200, getUrl)
	//c.Redirect(http.StatusMovedPermanently, getUrl)
}

// Print godoc
// @Summary 打印最近的一条短链接
// @Description
// @Tags url
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "打印短链接信息"
// @Router /print [get]
func Print(c *gin.Context) {
	shorturl := dao.Print()
	c.JSON(200, gin.H{
		"message": shorturl,
	})
}

