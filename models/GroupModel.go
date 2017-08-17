package models

import (
	"github.com/astaxie/beego"
)

/**
 *@project admin
 *@author fanceyh <fanceyh@qq.com>
 *@datetime 2017/7/13 11:38
 *@version 1.0
 */

//分组表
type Group struct {
	Id     int64
	Name   string  `orm:"size(100)" form:"Name"  valid:"Required"`
	Title  string  `orm:"size(100)" form:"Title"  valid:"Required"`
	Status int     `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	Sort   int     `orm:"default(1)" form:"Sort" valid:"Numeric"`
	Nodes  []*Node `orm:"reverse(many)"`
}

//自定义表名
func (g *Group) TableName() string {
	return beego.AppConfig.String("rbac_group_table")
}

// 设置引擎为 INNODB
func (g *Group) TableEngine() string {
	return "INNODB"
}
