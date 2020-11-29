package logic

import "github.com/gin-gonic/gin"

func Shorturl(c *gin.Context) {
	url := c.Query("url")

	c.JSON(200, gin.H{
		"message": url,
	})
}
