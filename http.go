package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/img", "./public/img")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "币图",
		})
	})

	r.Run(":80")// listen and serve on 0.0.0.0:8080
}
