package main

import (
	"github.com/gogf/gf-demos/v2/internal/cmd"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx = gctx.New()
	)
	if err := cmd.Main.Run(ctx); err != nil {
		g.Log().Fatal(ctx, err)
	}
}
