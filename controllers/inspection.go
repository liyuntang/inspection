package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myproject/models"
)

type InspectionController struct {
	beego.Controller
}

func (i *InspectionController)Get()  {
	product := i.GetString("product")
	clusterName := i.GetString("cluster_name")
	fmt.Println("product is", product, "clusterName is", clusterName)
	i.Data["data"] = models.Inspection(product, clusterName)
	i.TplName="result.html"
}
