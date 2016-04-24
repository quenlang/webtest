package controllers

import (
	"github.com/astaxie/beego"
	//"strings"
	//"golang.org/x/net/websocket"
	"webtest/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "login.html"
	return
}

func (this *MainController) Post() {
	vmAddr := "127.0.0.1:22"
	beego.Info(vmAddr)
	uname := this.Input().Get("user_name")
	beego.Info(uname)
	passwd := this.Input().Get("user_pwd")
	beego.Info(passwd)

	if uname == beego.AppConfig.String("user") && passwd == beego.AppConfig.String("password") {
		sh := &models.SSH{
			User: uname,
			Pwd:  passwd,
			Addr: vmAddr,
		}
		sh, err := sh.Connect()
		if err != nil {
			beego.Error(err)
		}

		wsAddr := "ws://" + this.Ctx.Request.Host + "/console/sshws/" + vmAddr
		beego.Info(wsAddr)
		this.Data["Uname"] = uname
		this.Data["WsAddr"] = wsAddr
		this.TplName = "console_main.html"
		return

	}
}
