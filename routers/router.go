package routers

import (
	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"
	"webtest/controllers"
	"webtest/models"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/console/login", &controllers.MainController{}, "post:Post")
	beego.Handler("/console/sshws/:vm_info", websocket.Handler(models.SSHWebSocketHandler))
}
