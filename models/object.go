package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	Videoslist []*Info
)

type Info struct {
	Id int            `orm:"column(id);auto"`
	Title string       `orm:"column(title);size(128)"`
	ImageHref string   `orm:"column(imagehref);size(128)"`
	Href string        `orm:"column(href);size(128)"`
	CreateDate time.Time  `orm:"column(createdate);type(datetime)"`
}


func AddOne(info *Info) (id int64,err error) {
	o := orm.NewOrm()
	id,err = o.Insert(info)
	fmt.Printf("id %v,err: %v",id,err)
	return id,err
}

func GetOne(id int) (info *Info,err error) {
	orm.Debug = true
	o := orm.NewOrm()
	v := &Info{Id:id}
	if err := o.Read(v);err == nil{
		return v,err
	}
	return nil,err

}

func GetAll() ([]*Info){
	orm.Debug = true
	o := orm.NewOrm()
	q := o.QueryTable("info")
	q.All(&Videoslist)
	return Videoslist

}

func Update(info *Info) (err error){
	o := orm.NewOrm()
	v := Info{Id: info.Id}
	if err = o.Read(&v);err == nil{
		if num,err := o.Update(info);err == nil{
			fmt.Println("Number of records updated in database:",num)
		}
	}
	return
}

func Delete(id int) (err error){
	o := orm.NewOrm()
	v := &Info{Id:id}
	if err := o.Read(v);err == nil{
		if num,err := o.Delete(v);err == nil{
			fmt.Println("Number of records deleted in database:",num)
		}
	}
	return
}
