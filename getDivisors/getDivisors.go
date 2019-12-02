package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type REQUEST struct {
	NUMBER int   `json:"number" binding:"required"`
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
		v1.POST("/", getDivisors)
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

// getDivisors returns all valid divisors for a number.
func getDivisors(context *gin.Context) {
	var req REQUEST
	if err := context.BindJSON(&req); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}
	for i := 1; i <= req.NUMBER; i++ {
		if req.NUMBER%i == 0 {
			req.RESULT = append(req.RESULT, i)
		}
	}
	context.AbortWithStatusJSON(http.StatusOK, req)
	return
}
