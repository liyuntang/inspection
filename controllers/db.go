package controllers

import "github.com/astaxie/beego"

type DbController struct {
	beego.Controller
}

func (db *DbController)Get()  {
	db.TplName="db.html"
}
