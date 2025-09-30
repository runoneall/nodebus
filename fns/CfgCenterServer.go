package fns

import (
	"fmt"
	"net"
	"nodebus/cli"
	"nodebus/configm"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func CfgCenterServer(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		ctx.Header("Pragma", "no-cache")
		ctx.Header("Expires", "0")
		ctx.Next()
	})

	r.GET("/status", func(ctx *gin.Context) {
		ctx.String(200, "running")
	})

	r.GET("/fresh", func(ctx *gin.Context) {
		switch err := configm.GetManager().LoadJSON(); err {
		case nil:
			ctx.String(200, "success")
		default:
			ctx.String(500, "%v", err)
		}
	})

	r.GET("/", func(ctx *gin.Context) {
		authHeader, ok := ctx.Request.Header["Auth"]
		if !ok || !slices.Contains(authHeader, *cli.CfgCenterAuth) {
			ctx.String(401, "unauth")
			return
		}

		ctx.JSON(200, configm.GetManager().ItemGetAll())
	})

	addr := net.JoinHostPort(
		*cli.CfgCenterHost,
		*cli.CfgCenterPort,
	)
	fmt.Println("监听", addr)

	if err := r.Run(addr); err != nil {
		panic(fmt.Errorf("不能启动服务: %v", err))
	}
}
