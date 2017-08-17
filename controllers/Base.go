package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
	m "admin/models"
	. "admin/lib"
	"errors"
	"fmt"
)

/**
 *@project admin
 *@author fanceyh <fanceyh@qq.com>
 *@datetime 2017/7/10 16:19
 *@version 1.0
 */

type BaseController struct {
	beego.Controller
}

type Tree struct {
	Id         int64      `json:"id"`
	Text       string     `json:"text"`
	IconCls    string     `json:"iconCls"`
	Checked    string     `json:"checked"`
	State      string     `json:"state"`
	Children   []Tree     `json:"children"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Url   string `json:"url"`
	Price int64  `json:"price"`
}

////确定是否需要验证
//func CheckAccess(params []string) bool {
//	if len(params) < 3 {
//		return false
//	}
//	for _, nap := range strings.Split(beego.AppConfig.String("not_auth_package"), ",") {
//		if params[1] == nap {
//			return false
//		}
//	}
//	return true
//}
//
////验证是否有权限
//func AccessDecision(params []string, accesslist map[string]bool) bool {
//	if CheckAccess(params) {
//		s := fmt.Sprintf("%s/%s/%s", params[1], params[2], params[3])
//		if len(accesslist) < 1 {
//			return false
//		}
//		_, ok := accesslist[s]
//		if ok != false {
//			return true
//		}
//	} else {
//		return true
//	}
//	return false
//}
//
////定义权限节点
//type AccessNode struct {
//	Id        int64
//	Name      string
//	Childrens []*AccessNode
//}
//
////权限列表
//func GetAccessList(uid int64) (map[string]bool, error) {
//	list, err := m.AccessList(uid)
//	if err != nil {
//		return nil, err
//	}
//	alist := make([]*AccessNode, 0)
//	for _, l := range list {
//		if l["Pid"].(int64) == 0 && l["Level"].(int64) == 1 {
//			anode := new(AccessNode)
//			anode.Id = l["Id"].(int64)
//			anode.Name = l["Name"].(string)
//			alist = append(alist, anode)
//		}
//	}
//	for _, l := range list {
//		if l["Level"].(int64) == 2 {
//			for _, an := range alist {
//				if an.Id == l["Pid"].(int64) {
//					anode := new(AccessNode)
//					anode.Id = l["Id"].(int64)
//					anode.Name = l["Name"].(string)
//					an.Childrens = append(an.Childrens, anode)
//				}
//			}
//		}
//	}
//	for _, l := range list {
//		if l["Level"].(int64) == 3 {
//			for _, an := range alist {
//				for _, an1 := range an.Childrens {
//					if an1.Id == l["Pid"].(int64) {
//						anode := new(AccessNode)
//						anode.Id = l["Id"].(int64)
//						anode.Name = l["Name"].(string)
//						an1.Childrens = append(an1.Childrens, anode)
//					}
//				}
//
//			}
//		}
//	}
//	accesslist := make(map[string]bool)
//	for _, v := range alist {
//		for _, v1 := range v.Childrens {
//			for _, v2 := range v1.Childrens {
//				vname := strings.Split(v.Name, "/")
//				v1name := strings.Split(v1.Name, "/")
//				v2name := strings.Split(v2.Name, "/")
//				str := fmt.Sprintf("%s/%s/%s", strings.ToLower(vname[0]), strings.ToLower(v1name[0]), strings.ToLower(v2name[0]))
//				accesslist[str] = true
//			}
//		}
//	}
//	return accesslist, nil
//}

//验证登录
func CheckLogin(username string, password string) (user m.User, err error) {
	user = m.GetUserByUsername(username)
	fmt.Println(user)
	if user.Id == 0 {
		return user, errors.New("用户不存在")
	}
	if user.Password != Pwdhash(password) {
		return user, errors.New("密码错误")
	}
	return user, nil
}

//判断是不是post
func (this *BaseController) IsPost() bool {
	return this.Ctx.Request.Method == "POST"
}

//json 基本返回
func (this *BaseController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
}

func (this *BaseController) GetTree() []Tree {
	nodes, _ := m.GetNodeTree(0, 1)
	beego.Info(&nodes)
	tree := make([]Tree, len(nodes))
	for k, v := range nodes {
		beego.Info(v["Id"].(int64))
		tree[k].Id = v["Id"].(int64)
		tree[k].Text = v["Title"].(string)
		tree[k].IconCls = v["IconCls"].(string)
		children, _ := m.GetNodeTree(v["Id"].(int64), 2)
		tree[k].Children = make([]Tree, len(children))
		for k1, v1 := range children {
			tree[k].Children[k1].Id = v1["Id"].(int64)
			tree[k].Children[k1].Text = v1["Title"].(string)
			tree[k].Children[k1].Attributes.Url = "/" + v["Name"].(string) + "/" + v1["Name"].(string)
		}
	}
	return tree
}
