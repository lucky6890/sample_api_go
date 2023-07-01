package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucky6890/sample_api_go/api/helpers"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, helpers.GenerateBaseResponse("Test", true, 0))
}
