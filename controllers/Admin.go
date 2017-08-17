package controllers

import "github.com/astaxie/beego"

/**
 *@project admin
 *@author fanceyh <fanceyh@qq.com>
 *@datetime 2017/7/18 16:40
 *@version 1.0
 */
type AdminController struct {
	BaseController
}

func (this *AdminController) Admin() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
	}
	tree := this.GetTree()
	this.Data["Menu"] = tree
	this.TplName = "admin.html"
}

func (this *AdminController) Welcome() {
	this.TplName = "welcome.html"
}

func (this *AdminController) GetMenu() {
	tree := this.GetTree()
	this.Data["json"] = &tree
	this.ServeJSON()
	return
}
