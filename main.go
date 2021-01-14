package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	r.GET("", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{})
	})
	r.Run(":9999")
}
