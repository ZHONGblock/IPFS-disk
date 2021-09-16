package routers

import (
	"ipfs-file-server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/list", &controllers.MainController{}, "GET:GetFileList")
	beego.Router("/upload", &controllers.MainController{}, "POST:UploadFile")
	beego.Router("/share", &controllers.MainController{}, "POST:ShareFile")
}
