package controllers

import (
	"beeadmin/models"
	//"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	//判断登陆是否合法
	if c.Input().Get("exit") == "true" {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Redirect("/", 301)
		return
	}
	if !checkAccount(c.Ctx) {
		c.Redirect("/", 302)
		return
	}
	/************传入数据******************/
	pre_page := 10
	pa := 1
	pa, _ = strconv.Atoi(c.Input().Get("p"))
	res, err := models.GetUsers(pa, pre_page)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["paginator"] = res
		c.TplName = "admin.html"
	}
}
func (c *AdminController) Dowload() {
	c.Redirect("/admin", 304)
	return
}
func (c *AdminController) Message() {
	id := c.Input().Get("id")
	num, _ := strconv.Atoi(id)
	message := models.GetNewUsers(num)
	newus := strconv.Itoa(message)
	c.Ctx.WriteString(newus)
}
