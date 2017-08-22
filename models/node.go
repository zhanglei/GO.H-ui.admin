package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
 *@project admin
 *@author fanceyh <fanceyh@qq.com>
 *@datetime 2017/7/13 11:43
 *@version 1.0
 */

type Node struct {
	Id      int64
	Title   string  `orm:"size(100)" form:"Title"  valid:"Required" description:"标题"`
	Name    string  `orm:"size(100)" form:"Name"  valid:"Required" description:"名称"`
	Level   int     `orm:"default(1)" form:"Level"  valid:"Required" description:"等级"`
	Pid     int64   `form:"Pid"  valid:"Required" description:"上一级ID"`
	Remark  string  `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)" description:"备注"`
	IconCls string  `orm:"size(32)" form:"IconCls"  valid:"Required" description:"图标"`
	Status  int     `orm:"default(2)" form:"Status" valid:"Range(1,2)" description:"状态"`
	Group   *Group  `orm:"rel(fk)"`
	Role    []*Role `orm:"rel(m2m)"`
}

//自定义表名
func (n *Node) TableName() string {
	return beego.AppConfig.String("rbac_node_table")
}

// 设置引擎为 INNODB
func (n *Node) TableEngine() string {
	return "INNODB"
}

func GetNodeTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(Node)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}
