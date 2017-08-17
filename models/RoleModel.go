package models

/**
 *@project admin
 *@author fanceyh <fanceyh@qq.com>
 *@datetime 2017/7/13 11:44
 *@version 1.0
 */
import (
	"github.com/astaxie/beego"
)

//角色表
type Role struct {
	Id     int64
	Title  string  `orm:"size(100)" form:"Title"  valid:"Required" description:"标题"`
	Name   string  `orm:"size(100)" form:"Name"  valid:"Required" description:"名称"`
	Remark string  `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)" description:"备注"`
	Status int     `orm:"default(2)" form:"Status" valid:"Range(1,2)" description:"状态"`
	Node   []*Node `orm:"reverse(many)"`
	User   []*User `orm:"reverse(many)"`
}

//自定义表名
func (r *Role) TableName() string {
	return beego.AppConfig.String("rbac_role_table")
}

// 设置引擎为 INNODB
func (r *Role) TableEngine() string {
	return "INNODB"
}
