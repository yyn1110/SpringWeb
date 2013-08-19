package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type HeadInfo struct {
	WebIndex       string
	WebSecondIndex string
	WebThridIndex  string
	WebMoreIndex   string
}
type MoreContent struct {
	Title       string
	SubContentF string
	SubContentS string
	SubContentT string
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
		WebIndex:       "首页",
		WebSecondIndex: "关于我们",
		WebThridIndex:  "我们的目标",
		WebMoreIndex:   "更多",
	}
	this.Data["More1"] = &MoreContent{
		Title:       "导演",
		SubContentF: "1.1111",
		SubContentS: "2.1111",
		SubContentT: "3.1111",
	}
	this.Data["More2"] = &MoreContent{
		Title:       "导演",
		SubContentF: "1.2111",
		SubContentS: "2.2111",
		SubContentT: "3.2111",
	}
	this.TplNames = "index.tpl"

}
