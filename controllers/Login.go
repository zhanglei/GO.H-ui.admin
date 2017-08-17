package controllers

import (
	"github.com/dchest/captcha"
	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	if this.IsPost() {
		username := this.GetString("username")
		password := this.GetString("password")
		captcha_id, vcode := this.GetString("captcha_id"), this.GetString("vcode")
		checkVcode := captcha.VerifyString(captcha_id, vcode) //验证码校验
		if checkVcode != true {
			this.Rsp(false, "验证码错误")
			return
		}
		user, err := CheckLogin(username, password)
		if err == nil {
			this.SetSession("userinfo", user)
			//accesslist, _ := GetAccessList(user.Id)
			//this.SetSession("accesslist", accesslist)
			this.Rsp(true, "登录成功")
			return
		} else {
			this.Rsp(false, err.Error())
			return
		}
	}
	userinfo := this.GetSession("userinfo")
	if userinfo != nil {
		this.Ctx.Redirect(302, "/admin")
	}
	captchaId := captcha.NewLen(4) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.TplName = "login.html"
}

func (this *LoginController) Logout() {
	this.DelSession("userinfo")
	this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
}
