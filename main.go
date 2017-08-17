package main

import (
	_ "admin/routers"
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"github.com/astaxie/beego/orm"
	"admin/models"
	"os"
)

func init() {
	beego.Handler("/captcha/*.png", captcha.Server(100, 40)) //验证码
	//注册数据库
	models.RegisterDB()

}

func main() {
	//开启orm调试
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)
	initArgs()
	beego.Run()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
	}
}
