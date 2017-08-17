package models

import (
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"errors"
	. "admin/lib"
	"github.com/astaxie/beego/validation"
	"log"
)

/**
 *@project admin
 *@author fanceyh <fanceyh@qq.com>
 *@datetime 2017/7/13 11:45
 *@version 1.0
 */

//用户表
type User struct {
	Id            int64
	Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)" description:"用户名"`
	Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)" description:"密码"`
	Repassword    string    `orm:"-" form:"Repassword" valid:"Required" `
	Nickname      string    `orm:"unique;size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(2)" description:"昵称"`
	Phone         string    `orm:"default()" form:"Phone" valid:"Required;MaxSize(11);MinSize(11)" description:"手机号"`
	Email         string    `orm:"size(32)" form:"Email" valid:"Email" description:"邮箱"`
	Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)" description:"备注"`
	Status        int       `orm:"default(1)" form:"Status" valid:"Range(1,2)" description:"状态"`
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-" description:"最后登录时间"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" description:"创建时间"`
	IsDel         int       `orm:"default(2)" form:"IsDel" valid:"Range(1,2)" description:"是否删除"`
	Deltime       time.Time `orm:"null;type(datetime)" form:"-" description:"删除时间"`
	Role          []*Role   `orm:"rel(m2m)"`
}

//自定义表名
func (u *User) TableName() string {
	return beego.AppConfig.String("rbac_user_table")
}

// 设置引擎为 INNODB
func (u *User) TableEngine() string {
	return "INNODB"
}

//根据用户名获取用户信息
func GetUserByUsername(username string) (user User) {
	user = User{Username: username}
	o := orm.NewOrm()
	o.Read(&user, "Username")
	return user
}

//获取用户列表
func Getuserlist() (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	qs.Filter("is_del__exact", 0).Values(&users)
	count, _ = qs.Count()
	return users, count
}

//验证用户信息
func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

//添加用户
func AddUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := new(User)
	user.Username = u.Username
	user.Password = Strtomd5(u.Password)
	user.Nickname = u.Nickname
	user.Phone = u.Phone
	user.Email = u.Email
	user.Remark = u.Remark
	user.Status = u.Status

	id, err := o.Insert(user)
	return id, err
}

//更新用户
func UpdateUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := make(orm.Params)
	if len(u.Username) > 0 {
		user["Username"] = u.Username
	}
	if len(u.Nickname) > 0 {
		user["Nickname"] = u.Nickname
	}
	if len(u.Email) > 0 {
		user["Email"] = u.Email
	}
	if len(u.Remark) > 0 {
		user["Remark"] = u.Remark
	}
	if len(u.Password) > 0 {
		user["Password"] = Strtomd5(u.Password)
	}
	if u.Status != 0 {
		user["Status"] = u.Status
	}
	if u.IsDel != 0 {
		user["IsDel"] = u.IsDel
		if u.IsDel == 2 {
			user["deltime"] = nil
		} else {
			user["deltime"] = time.Now()
		}

	}
	if u.Id == 1 {
		if user["IsDel"] == 1 || user["Status"] == 2 {
			return 0, errors.New("无法操作！")
		}
	}

	if len(user) == 0 {
		return 0, errors.New("修改内容不能为空！")
	}
	var table User
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
	if num == 0 {
		return 0, errors.New("更新失败！")
	} else {
		return num, err
	}

}

//根据id删除用户
func DelUserById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&User{Id: Id})
	return status, err
}
