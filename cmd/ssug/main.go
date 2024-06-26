package ssug

import (
	"fmt"
	"io"
	"ssug/internal/base"
	"ssug/internal/utils"
	"ssug/modules/api"
	base2 "ssug/modules/base"
	"ssug/modules/data"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	base.PrintBanner()
}

func Main() {
	p := base.ParsePara()
	data.Redirect.Init(p.Key, p.TTL)
	base2.RemoveExp()
	base2.SetInitLen(p.InitLen)

	go base2.AutoRemove()

	g := gin.Default()
	g.POST("/add", api.AddMapping)
	g.GET("/:short", api.GetMapping)
	g.GET("/", api.Happy)
	g.POST("/", api.Happy)

	utils.Logger.Info(fmt.Sprintf("在%s上启动短链服务...", p.HostPort))
	r := g.Run(p.HostPort)
	if r != nil {
		utils.Logger.Fatal(r)
	}
}
