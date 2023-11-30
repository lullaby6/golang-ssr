package main

import (
	"fmt"
	"main/src/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	pages := utils.GetFilesFromDirWithSuffix("./src/pages", ".html")
	components := utils.GetFilesFromDirWithSuffix("./src/components", ".html")
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

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "My Title",
			"Content": "Hello world!",
			"Texts":   texts,
		})
	})

	router.NoRoute(func(ctx *gin.Context) {
		request := strings.Replace(ctx.Request.URL.Path, "/", "", 1)

		file := fmt.Sprintf("./src/pages/%s/index.html", request)
		if _, err := os.Stat(file); err == nil {
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.File(file)
			return
		}

		file = fmt.Sprintf("./src/pages/%s.html", request)
		if _, err := os.Stat(file); err == nil {
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.File(file)
			return
		}

		ctx.Writer.WriteHeader(http.StatusNotFound)
		ctx.File("./src/pages/404.html")
	})

	PORT := utils.GetEnvOrDefault("PORT", "3000")
	HOST := utils.GetEnvOrDefault("HOST", "127.0.0.1")
	router.Run(fmt.Sprintf("%s:%s", HOST, PORT))
}
