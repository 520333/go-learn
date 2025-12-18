package main

import (
	_ "beegoTest/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
