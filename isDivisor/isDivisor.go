package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DIVISION struct {
	DIVISOR  int  `json:"divisor" binding:"required"`
	DIVIDEND int  `json:"dividend" binding:"required"`
	RESULT   bool `json:"result"`
}

func main() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1/")
	{
		v1.POST("/", isDivisor)
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

// isDivisor checks if a number is valid divisor.
func isDivisor(context *gin.Context) {
	var division DIVISION
	if err := context.BindJSON(&division); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if division.DIVIDEND%division.DIVISOR != 0 {
		division.RESULT = false
		context.AbortWithStatusJSON(http.StatusOK, division)
		return
	}
	division.RESULT = true
	context.AbortWithStatusJSON(http.StatusOK, division)
}
