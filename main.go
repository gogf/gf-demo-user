package main

import (
	_ "gitee.com/johng/gf-demos/boot"
	_ "gitee.com/johng/gf-demos/router"
	"gitee.com/johng/gf/g"
)

func main() {
	g.Server().Run()
}
