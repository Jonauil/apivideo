package models

import (
	"errors"
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	Objects map[string]*Object
	Videos  map[string]*Info
	num int64
	//Videoslist map[string] []*Info
	VideosMaps []orm.Params
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
	//Videoslist := make(map[string] []*Info)
	Objects["hjkhsbnmn123"] = &Object{"hjkhsbnmn123", 100, "astaxie"}
	Objects["mjjkxsxsaa23"] = &Object{"mjjkxsxsaa23", 101, "someone"}
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

func GetAll() (map[string] interface{}) {
	//info := Info{}
	/*
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
	}*/
	o := orm.NewOrm()
    videoslist := make(map[string] interface{})
   // VideoData := make(map[string] interface{})
	num,err := o.Raw("SELECT * FROM info").Values(&VideosMaps)
	fmt.Println("num:",num)
	for k,v := range VideosMaps{
		fmt.Println("key:",k,"value:",v)
		videoslist["videodata"] = v

	}
	if err != nil {
		fmt.Println("查询不到")
	}
	//Videos["videolist"] = &info
	return videoslist
}

func Update(ObjectId string, Score int64) (err error) {
	if v, ok := Objects[ObjectId]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ObjectId Not Exist")
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
