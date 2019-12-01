package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DIVISION struct {
	DIVISOR  int  `json:"divisor" binding:"required"`
	DIVIDEND int  `json:"dividend" binding:"required"`
	RESULT   bool `json:"result" binding:"required"`
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/")
	{
		v1.POST("/", isDivisor)
	}
	router.Run()
}

// isDivisor checks if a number is valid divisor.
func isDivisor(context *gin.Context) {
	var division DIVISION
	var err error
	dividend, _ := context.GetQuery("dividend")
	divisor, _ := context.GetQuery("divisor")
	division.DIVIDEND, err = strconv.Atoi(dividend)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}
	division.DIVISOR, err = strconv.Atoi(divisor)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}
	if division.DIVIDEND%division.DIVISOR != 0 {
		division.RESULT = false
		context.JSON(http.StatusOK, division)
	}
	division.RESULT = true
	context.JSON(http.StatusOK, division)
}
