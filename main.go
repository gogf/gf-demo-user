package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "github.com/gogf/gf-demo-user/v2/internal/logic"

	"github.com/gogf/gf-demo-user/v2/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
