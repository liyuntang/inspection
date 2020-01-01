package routers

import (
	"myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/db",  &controllers.DbController{})
    beego.Router("/inspection", &controllers.InspectionController{})
	beego.Router("/inspectionAll", &controllers.InspectionAllController{})
}
