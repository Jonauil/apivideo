package controllers

import (
	"apivideo/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"strconv"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *ObjectController) Post() {
	var ob models.Info
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if id,err := models.AddOne(&ob);err == nil{
		o.Data["json"] = map[string]int64{"id":id}
	}else{
		o.Data["json"] = err.Error()
	}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *ObjectController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		videoObject,_ :=strconv.Atoi(objectId)
		ob, err := models.GetOne(videoObject)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ObjectController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

func (o *ObjectController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	id,_ := strconv.Atoi(objectId)
	info := models.Info{Id:id}
	json.Unmarshal(o.Ctx.Input.RequestBody, &info)

	if err := models.Update(&info); err == nil {
		o.Data["json"] = "update success!"
	}else{
		o.Data["json"] = err.Error()
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *ObjectController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	id,_ := strconv.Atoi(objectId)
	if err := models.Delete(id);err == nil{
		o.Data["json"] = "delete success!"
	}else{
		o.Data["json"] = err.Error()
	}
	o.ServeJSON()
}

