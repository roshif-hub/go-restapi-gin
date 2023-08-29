package main

import (
	"github.com/gin-gonic/gin"
	"github.com/roshif-study/go-restapi-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/product", produccontroller.Index)
	r.GET("/api/product/:id", produccontroller.Show)
	r.POST("/api/product", produccontroller.Create)
	r.PUT("/api/product/:id", produccontroller.Update)
	r.DELETE("/api/product", produccontroller.Delete)

	r.Run()
}