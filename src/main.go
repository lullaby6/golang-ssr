package main

import (
	"fmt"
	"main/src/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()

	pages := utils.GetHTMLFilesFromDir("./src/pages")
	components := utils.GetHTMLFilesFromDir("./src/components")
	files := append(pages, components...)
	r.LoadHTMLFiles(files...)

	r.Use(utils.Cors("*", "POST,HEAD,PATCH,OPTIONS,GET,PUT"))

	r.Static("/assets", "./sc/assets")

	r.GET("/", func(c *gin.Context) {
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

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "My Title",
			"Content": "Hello world!",
			"Texts":   texts,
		})
	})

	r.NoRoute(func(c *gin.Context) {
		request := strings.Replace(c.Request.URL.Path, "/", "", 1)

		for _, path := range pages {
			file := strings.Replace(path, "src\\pages\\", "", 1)
			if request+".html" == file {
				c.File(path)
				return
			}
		}

		c.Writer.WriteHeader(http.StatusNotFound)
		c.File("./src/pages/404.html")
	})

	PORT := utils.GetEnvOrDefault("PORT", "3000")
	HOST := utils.GetEnvOrDefault("HOST", "127.0.0.1")
	r.Run(fmt.Sprintf("%s:%s", HOST, PORT))
}
