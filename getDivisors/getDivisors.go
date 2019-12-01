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
	router := gin.Default()
	v1 := router.Group("/api/v1/")
	{
		v1.POST("/", getDivisors)
	}
	router.Run()
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
