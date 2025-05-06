package main

import "github.com/gin-gonic/gin"

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})

	r.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "json",
			"msg":     "hello world",
		})
	})

	r.GET("/jsonArticle", func(c *gin.Context) {
		a := Article{
			Title:   "基督山伯爵",
			Desc:    "基督山伯爵",
			Content: "基督山伯爵",
		}
		c.JSON(200, a)
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{
			"message": "xml",
			"msg":     "hello world",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.HTML(200, "news.html", gin.H{
			"title": "news",
		})
	})

	r.GET("/content", func(c *gin.Context) {

		news := &Article{
			Title:   "基督山伯爵",
			Content: "基督山伯爵",
		}

		c.HTML(200, "content.html", gin.H{
			"title": "基督山伯爵",
			"news":  news,
		})
	})

	r.Run() // listen and serve on

}
