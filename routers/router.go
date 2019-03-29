package routers

import (
	"BeegoSourceAnalysis/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/score", &controllers.Model{})
}
