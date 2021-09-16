package controllers

import (
	"encoding/json"
	"ipfs-file-server/models"
	"log"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "web.html"
}

func (c *MainController) GetFileList() {
	list := models.GetFileList()

	c.Data["json"] = list
	c.ServeJSON()
}

func (c *MainController) ShareFile() {
	var req models.SharedReq
	c.postReqJsonAnly(&req)

	err := models.ShareFile(req.FileCid, req.Shared)
	c.responseErr(err)
}

func (c *MainController) UploadFile() {
	f, h, _ := c.GetFile("srcfile")
	f.Close()

	name := h.Filename
	size := h.Size

	file, _, err := c.Ctx.Request.FormFile("srcfile")
	if err != nil {
		log.Println("ipfs add file error:", err)
		c.Ctx.ResponseWriter.WriteHeader(403)
		return
	}
	defer file.Close()

	err = models.SaveFileToIpfs(name, size, file)
	if err != nil {
		log.Println("ipfs add file error:", err)
		c.Ctx.ResponseWriter.WriteHeader(403)
		return
	}

	c.TplName = "web.html"
}

func (c *MainController) postReqJsonAnly(req interface{}) {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if nil != err {
		c.Ctx.ResponseWriter.WriteHeader(403)
		resp := &models.RespMsg{Msg: err.Error()}
		c.Data["json"] = resp
		c.ServeJSON()
	}
}

func (c *MainController) responseErr(err error) {
	if nil == err {
		c.Data["json"] = &models.RespMsg{Msg: "ok"}
	} else {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = &models.RespMsg{Msg: err.Error()}
	}
	c.ServeJSON()
}
