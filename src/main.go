package main

import (
	"fmt"
	"main/src/components"
	"main/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	router.Static("/assets", "./src/assets")

	router.GET("/", func(ctx *gin.Context) {
		texts := []string{
			"Text1",
			"Text2",
			"Text3",
		}

		textsComponent := utils.JoinStringWithCallback(texts, components.Text)

		props := map[string]interface{}{
			"Title":   "My Title",
			"Content": "Hello world!",
			"Texts":   textsComponent,
		}

		html, err := utils.ParseHTML("./src/pages/index.html", props)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", html)
	})

	router.NoRoute(func(ctx *gin.Context) {
		html, err := utils.GetHTML("./src/pages/404.html")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", html)
	})

	PORT := utils.GetEnvOrDefault("PORT", "3000")
	HOST := utils.GetEnvOrDefault("HOST", "127.0.0.1")
	router.Run(fmt.Sprintf("%s:%s", HOST, PORT))
}
