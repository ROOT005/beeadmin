package controllers

import (
	"beeadmin/models"
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	users, err := models.GetAllUsers()
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Users"] = users
	}
	c.Data["Pre"] = "disabled"
	c.Data["Next"] = ""
	if !checkAccount(c.Ctx) {
		c.Redirect("/", 302)
		return
	}
	if c.Input().Get("exit") == "true" {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Redirect("/", 301)
		return
	}
	c.TplName = "admin.html"
}
