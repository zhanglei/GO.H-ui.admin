package controllers

import (
	m "admin/models"
)

/**
 *@project admin
 *@author fanceyh <fanceyh@qq.com>
 *@datetime 2017/8/14 15:54
 *@version 1.0
 */

type AdminuserController struct {
	BaseController
}

func (this *AdminuserController) Index() {
	//page, _ := this.GetInt64("page")
	//page_size, _ := this.GetInt64("rows")
	//sort := this.GetString("sort")
	//order := this.GetString("order")
	//if len(order) > 0 {
	//	if order == "desc" {
	//		sort = "-" + sort
	//	}
	//} else {
	//	sort = "Id"
	//}
	users, count := m.Getuserlist()
	this.Data["users"] = &users
	this.Data["count"] = &count
	this.TplName = "rbac/member-list.html"
}

func (this *AdminuserController) UpdateUser() {
	u := m.User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	num, err := m.UpdateUser(&u)
	if err == nil && num > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *AdminuserController) DelUser() {
	u := m.User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	num, err := m.UpdateUser(&u)
	if err == nil && num > 0 {
		this.Rsp(true, "成功！")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}
