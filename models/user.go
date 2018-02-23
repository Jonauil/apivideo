package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	UserList map[string]*User
	Userlist []*UserInfo

)

type User struct {
	Id       string
	Username string
	Password string
}
type UserInfo struct {
	Uid int `orm:"column(uid);auto"`
	Username string `orm:"column(username);size(128)"`
	Userphoto string `orm:"column(userphoto);size(128)"`
}


func AddUser(uInfo *UserInfo) (int int64,err error) {
	o := orm.NewOrm()
	id,err := o.Insert(uInfo)
	fmt.Printf("id: %v,err: %v",id,err)
	return id,err
}

func GetUser(uid int)(uinfo *UserInfo,err error){
	orm.Debug = true
	o := orm.NewOrm()
	userinfo := &UserInfo{Uid:uid}
	if err := o.Read(userinfo);err == nil {
		return userinfo,nil
	}
	return nil,err
}

func GetAllUsers() []*UserInfo{
	orm.Debug = true
	o := orm.NewOrm()
	q := o.QueryTable("user_info")
	q.All(&Userlist)
	return Userlist
}

func UpdateUser(info *UserInfo)(err error){
	o := orm.NewOrm()
	vinfo := UserInfo{Uid:info.Uid}
	if o.Read(&vinfo) == nil {
		if num,err := o.Update(info);err == nil{
			fmt.Println("Number of records updated in database",num)
		}
	}
	return
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid int) (err error){
	orm.Debug = true
	o := orm.NewOrm()
	user := &UserInfo{Uid:uid}
	if err := o.Read(user);err == nil{
		if num,err := o.Delete(user); err == nil {
			fmt.Println("Number of records deleted in database:",num)
		}
	}
	return
}
