package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type REQUEST struct {
	X      []int `json:"x" binding:"required"`
	RESULT int   `json:"result"`
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/")
	{
		v1.POST("/", maximum)
	}
	router.Run()
}

func maximum(context *gin.Context) {
	var req REQUEST
	if err := context.BindJSON(&req); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}
	for i, e := range req.X {
		if i == 0 || e > req.RESULT {
			req.RESULT = e
		}
	}
	context.AbortWithStatusJSON(http.StatusOK, req)
	return
}
