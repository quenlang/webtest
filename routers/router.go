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
	beego.Handler("/console/sshws/:127.0.0.1:22", websocket.Handler(models.SSHWebSocketHandler))
}
