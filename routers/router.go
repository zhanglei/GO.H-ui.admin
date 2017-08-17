package routers

import (
	"admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "*:Login")     //登录
	beego.Router("/logout", &controllers.LoginController{}, "*:Logout")   //退出登录
	beego.Router("/admin", &controllers.AdminController{}, "*:Admin")     //后台
	beego.Router("/welcome", &controllers.AdminController{}, "*:Welcome") //后台首页
	beego.Router("/getmenu", &controllers.AdminController{}, "*:GetMenu") //获取菜单

	//RBAC
	beego.Router("/rbac/user/index", &controllers.AdminuserController{}, "*:Index")           //用户管理
	beego.Router("/rbac/user/UpdateUser", &controllers.AdminuserController{}, "*:UpdateUser") //用户修改
	beego.Router("/rbac/user/DelUser", &controllers.AdminuserController{}, "*:DelUser")       //用户删除
	//demo
	beego.Router("/captcha", &controllers.CaptchaController{}) //验证码
}
