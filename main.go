package main

import (
	_ "ipfs-file-server/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
