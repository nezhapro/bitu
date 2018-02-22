package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
    "strings"
)

func main() {
	r:= gin.Default()



	//r.LoadHTMLGlob("templates/*")
  r.LoadHTMLGlob("templates/*")
	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/img", "./public/img")

	r.GET("/", func(c *gin.Context) {
    if (strings.Contains(c.Request.Host, "bitu.fun")){
      c.HTML(http.StatusOK, "bitu.tmpl", gin.H{
  			"title": "币图",
  		})
    } else if(strings.Contains(c.Request.Host, "nezha.pro")){
      c.HTML(http.StatusOK, "nezha.tmpl", gin.H{
        "title": "哪吒",
      })
    } else if(strings.Contains(c.Request.Host, "bibang.fun")){
      c.HTML(http.StatusOK, "bibang.tmpl", gin.H{
        "title": "币帮",
      })
    }else if(strings.Contains(c.Request.Host, "jianbi.fun")){
      c.HTML(http.StatusOK, "jianbi.tmpl", gin.H{
        "title": "鉴币",
      })
    }else if(strings.Contains(c.Request.Host, "kuaiwa.fun")){
      c.HTML(http.StatusOK, "kuaiwa.tmpl", gin.H{
        "title": "块蛙",
      })
    }else if(strings.Contains(c.Request.Host, "qkc.fun")){
      c.HTML(http.StatusOK, "qkc.tmpl", gin.H{
        "title": "区块车",
      })
    }else{
      c.HTML(http.StatusOK, "nezha.tmpl", gin.H{
        "title": "",
      })
    }
	})

  r.GET("/about", func(c *gin.Context) {
    fmt.Println(c.Request)
    //c.String(200,"ok")
    c.JSON(http.StatusOK, gin.H{"status": c.Request.Host})
  })



	r.Run(":80")// listen and serve on 0.0.0.0:8080
}
