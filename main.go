package main

import (
    "gitee.com/johng/gf/g"
    _ "gitee.com/johng/gf-cases/boot"
    _ "gitee.com/johng/gf-cases/router"
)

func main() {
    g.Server().Run()
}