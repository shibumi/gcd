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
	router := gin.Default()
	v1 := router.Group("/api/v1/")
	{
		v1.POST("/", commonNumbers)
	}
	router.Run()
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
