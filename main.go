package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})

	// 處理 404 請求
	r.NoRoute(func(c *gin.Context) {
		// Accept 訊息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString , "text/html") {
			c.String(http.StatusNotFound,"404 頁面返回")
		}else {
			c.JSON(http.StatusNotFound,gin.H{
				"error_code" : http.StatusNotFound,
				"error_message" : "路由未定義，請確認!!",
			})
		}

	})

	r.Run(":8000")
}
