package main

import (
	"fmt"
	"go-goof/handlers"

	"github.com/gin-gonic/gin"
	"github.com/gogs/gogs/pkg/tool"
	"github.com/pstember/go-goof/hello"
)

func main() {
	fmt.Println(tool.MD5(hello.Hello()))
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.Run("0.0.0.0:8082") // listen and serve on 0.0.0.0:8080
}
