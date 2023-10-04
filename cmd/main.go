package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load env : %v", err)
	}

	token := os.Getenv("SLACK_TOKEN")

	fmt.Printf("This is Slack Token: %v \n", token)

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
