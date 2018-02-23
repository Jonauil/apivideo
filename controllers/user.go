package controllers

import (
	"apivideo/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"strconv"
)

var (
	userInfo models.UserInfo
)
// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	json.Unmarshal(u.Ctx.Input.RequestBody,&userInfo)
	if uid,err := models.AddUser(&userInfo);err == nil {
          u.Data["json"] = map[string]int64{"uid":uid}
	}else{
		  u.Data["json"] = err
	}
    u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	objectId := u.Ctx.Input.Param(":uid")
	if objectId != "" {
		userObject,_ := strconv.Atoi(objectId)
		user, err := models.GetUser(userObject)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	objectId := u.Ctx.Input.Param(":uid")
	uid,_ := strconv.Atoi(objectId)
	uinfo := models.UserInfo{Uid:uid}
	json.Unmarshal(u.Ctx.Input.RequestBody,&uinfo)

	if err := models.UpdateUser(&uinfo);err == nil {
		u.Data["json"] = "update success!"
	}else{
		u.Data["json"] = err.Error()
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.Ctx.Input.Param(":uid")
	id,_ := strconv.Atoi(uid)
	if err := models.DeleteUser(id);err == nil{
		u.Data["json"] = "delete success!"
	}else{
		u.Data["json"] = err.Error()
	}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

