package main

import (
	_ "gitee.com/johng/gf-cases/boot"
	_ "gitee.com/johng/gf-cases/router"
	"gitee.com/johng/gf/g"
)

func main() {
	g.Server().Run()
}
