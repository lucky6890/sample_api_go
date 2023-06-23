package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test": "test",
	})
	return
}
