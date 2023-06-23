package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucky6890/sample_api_go/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	handler := handlers.NewTestHandler()

	r.GET("/", handler.Test)
}
