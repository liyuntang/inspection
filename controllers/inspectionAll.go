package controllers

import (
	"github.com/astaxie/beego"
	"myproject/models"
)

type InspectionAllController struct {
	beego.Controller
}

func (i *InspectionAllController)Get() {
	i.Data["data"] = models.InspectionAll()
	i.TplName= "result.html"
}
