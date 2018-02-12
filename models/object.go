package models

import (
	"errors"
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	Objects map[string]*Object
	Videos map[string]*Info
)

type Object struct {
	ObjectId   string
	Score      int64
	PlayerName string
}

type Info struct {
	Id int            `orm:"column(id);auto"`
	Title string       `orm:"column(title);size(128)"`
	ImageHref string   `orm:"column(imagehref);size(128)"`
	Href string        `orm:"column(href);size(128)"`
	CreateDate time.Time  `orm:"column(createdate);type(datetime)"`
}

func init() {
	Objects = make(map[string]*Object)
	Videos = make(map[string]*Info)
	Objects["hjkhsbnmn123"] = &Object{"hjkhsbnmn123", 100, "astaxie"}
	Objects["mjjkxsxsaa23"] = &Object{"mjjkxsxsaa23", 101, "someone"}
}

func AddOne(object Object) (ObjectId string) {
	object.ObjectId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Objects[object.ObjectId] = &object
	return object.ObjectId
}

/*func GetOne(ObjectId string) (object *Object, err error) {
	if v, ok := Objects[ObjectId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}*/
func GetOne(id int) (map[string]*Info) {
/*	info := Info{Id:id}
	fmt.Println("info====",info)
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	errvideo := o.Read(&info)
	if errvideo != nil{
		fmt.Println("errvideo",errvideo)
	}
	Videos["videolist"] = &info
	*//*if v, ok := Objects[ObjectId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")*//*
	return Videos*/
	info := Info{Id:id}
	fmt.Println("info:",info)
	orm.Debug = true
	o := orm.NewOrm()
	err := o.Read(&info)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	}else if err == orm.ErrMissPK{
		fmt.Println("查询不到主键")
	}else {
		fmt.Printf("Id %v \n title %v ImageHref %v Href %v CreateDate %v",info.Id,info.Title,info.ImageHref,info.Href,info.CreateDate)
		Videos["videllist"] = &info
	}
	return Videos

}

//func GetAll() map[string]*Object {
//	return Objects
//}

func GetAll() map[string]*Info {
	info := Info{Id:1}
	fmt.Println("info=",info)
	orm.Debug = true
	o := orm.NewOrm()
	err := o.Read(&info)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	}else if err == orm.ErrMissPK{
		fmt.Println("查询不到主键")
	}else {
		fmt.Printf("Id %v \n title %v ImageHref %v Href %v CreateDate %v",info.Id,info.Title,info.ImageHref,info.Href,info.CreateDate)
		Videos["videllist"] = &info
	}
	return Videos
}

func Update(ObjectId string, Score int64) (err error) {
	if v, ok := Objects[ObjectId]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ObjectId Not Exist")
}

func Delete(ObjectId string) {
	delete(Objects, ObjectId)
}

