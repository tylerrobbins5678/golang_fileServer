package controller

import (
	"fileManager/file/service"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/*path", service.GetNames)
	r.POST("/*path", service.Create)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
