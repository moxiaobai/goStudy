package routers

import (
	"github.com/moxiaobai/goStudy/webdemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
