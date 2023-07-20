package main

import (
	"github.com/lucky6890/sample_api_go/api"
	"github.com/lucky6890/sample_api_go/config"
	"github.com/lucky6890/sample_api_go/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
}
