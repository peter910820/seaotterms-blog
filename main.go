package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	blog := gin.Default()
	blog.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "哈哈屁眼")
	})

	blog.Run(":8000")
}
