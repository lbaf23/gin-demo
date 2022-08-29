package main

import (
	"fmt"
	"gin-demo/conf"
	"gin-demo/models"
	"gin-demo/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.Init()
	gin.SetMode(conf.Config.RunMode)
	r := gin.Default()
	models.Init()

	r.Use(routers.CorsMiddleware())
	routers.Init(r)

	r.Run(fmt.Sprintf(":%s", conf.Config.HttpPort))
}
