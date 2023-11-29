package main

import (
	"fmt"
	"main/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	pages := utils.GetHTMLFilesFromDir("./src/pages")
	components := utils.GetHTMLFilesFromDir("./src/components")
	files := append(pages, components...)
	router.LoadHTMLFiles(files...)

	router.Use(utils.Cors("*", "POST,HEAD,PATCH,OPTIONS,GET,PUT"))

	router.Static("/assets", "./src/assets")

	router.GET("/", func(ctx *gin.Context) {
		texts := []map[string]interface{}{
			{
				"Title":   "Title 1",
				"Content": "Content 1",
			},
			{
				"Title":   "Title 2",
				"Content": "Content 2",
			},
			{
				"Title":   "Title 3",
				"Content": "Content 3",
			},
		}

		ctx.HTML(http.StatusNotFound, "index.html", gin.H{
			"Title":   "My Title",
			"Content": "Hello world!",
			"Texts":   texts,
		})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	PORT := utils.GetEnvOrDefault("PORT", "3000")
	HOST := utils.GetEnvOrDefault("HOST", "127.0.0.1")
	router.Run(fmt.Sprintf("%s:%s", HOST, PORT))
}
