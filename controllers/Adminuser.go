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

//后台用户管理列表
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
	this.TplName = "rbac/admin-list.html"
}

//更新后台用户信息
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

//删除后台用户
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

//添加用户
func (this *AdminuserController) AddUser() {
	if this.IsPost() {
		u := m.User{}
		if err := this.ParseForm(&u); err != nil {
			//handle error
			this.Rsp(false, err.Error())
			return
		}
		id, err := m.AddUser(&u)
		if err == nil && id > 0 {
			this.Rsp(true, "成功")
			return
		} else {
			this.Rsp(false, err.Error())
			return
		}
	} else {
		this.TplName = "rbac/admin-add.html"
	}
}

//修改用户信息页面
func (this *AdminuserController) EditUser() {
	id, _ := this.GetInt64("Id")
	user := m.GetUserById(id)
	this.Data["user"] = &user
	this.Data["id"] = id
	this.TplName = "rbac/admin-edit.html"
}
