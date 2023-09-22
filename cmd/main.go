package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./web/templates/*.html")

	// HTML 템플릿 파일을 렌더링하는 라우터
	router.GET("/", func(c *gin.Context) {
		// HTML 렌더링
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/greeting", func(c *gin.Context) {
		// HTML 렌더링
		c.String(http.StatusOK, "HELLO WORLD!")
	})

	router.POST("/result", func(c *gin.Context) {
		slackID := c.PostForm("slackID")
		c.HTML(http.StatusOK, "result.html", gin.H{"slackID": slackID})
	})

	router.Run(":8080")
}
