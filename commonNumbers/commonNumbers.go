package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type REQUEST struct {
	X      []int `json:"x" binding:"required"`
	Y      []int `json:"y" binding:"required"`
	RESULT []int `json:"result"`
}

func main() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1/")
	{
		v1.POST("/", commonNumbers)
	}
	basic := router.Group("/")
	{
		basic.GET("/ping", ping)
	}
	return router
}

func ping(context *gin.Context) {
	context.AbortWithStatus(http.StatusOK)
	return
}

// commonNumbers checks for common numbers in two int slices.
func commonNumbers(context *gin.Context) {
	var req REQUEST
	if err := context.BindJSON(&req); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	for _, xe := range req.X {
		for _, ye := range req.Y {
			if xe == ye {
				req.RESULT = append(req.RESULT, xe)
			}
		}
	}
	context.AbortWithStatusJSON(http.StatusOK, req)
	return
}
