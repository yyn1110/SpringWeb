package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type HeadInfo struct {
	WebName        string
	WebIndex       string
	WebSecondIndex string
	WebMoreIndex   string
}

var Version string
var AppName string

func init() {
	Version = "1.0"
	AppName = "大运传媒"
}
func (this *MainController) Get() {
	this.Data["version"] = Version
	this.Data["appname"] = AppName
	this.Data["indexImg1"] = "/static/img/slide-01.jpg"
	this.Data["indexImg2"] = "/static/img/slide-02.jpg"
	this.Data["indexImg3"] = "/static/img/slide-03.jpg"
	this.Data["headinfo"] = &HeadInfo{
		WebName:        "大运传媒",
		WebIndex:       "首页",
		WebSecondIndex: "关于",
		WebMoreIndex:   "更多",
	}
	this.TplNames = "index.tpl"

}
