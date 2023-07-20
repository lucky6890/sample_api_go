package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lucky6890/sample_api_go/api/routers"
	"github.com/lucky6890/sample_api_go/config"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	RegisterRoutes(r)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		panic(err)
	}
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("api")
	v1 := api.Group("/v1")
	{
		test := v1.Group("/test")
		routers.TestRouter(test)
	}
}
