package controllers

import (
	"beeadmin/models"
	"encoding/csv"
	"github.com/astaxie/beego"
	"os"
	"strconv"
	"time"
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

//下数据载客户
func (c *AdminController) Dowload() {
	filename := "data/" + time.Now().Format("2006-01-02") + ".xls"
	f, err := os.Create(filename)
	if err != nil {
		c.Ctx.WriteString(filename)
	}
	//从数据库取出数据
	users, err := models.GetAllUsers()
	if err != nil {
		c.Ctx.WriteString("下载失败!")
	}
	//导出数据csv格式
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") //声明编码格式为UTF-8
	w := csv.NewWriter(f)
	w.Write([]string{"姓名", "贷款金额/万", "联系电话", "时间"})
	for _, user := range users {
		w.Write([]string{user.Name, user.Account, user.PhoneNum, user.Created})
	}
	w.Flush()
	c.Ctx.Output.Download(filename)
}

//删除用户数据
func (c *AdminController) Delete() {
	sid := c.Input().Get("id")
	id, _ := strconv.Atoi(sid)
	err := models.DeleteUser(int64(id))
	if err != nil {
		c.Ctx.WriteString("删除失败!")
	} else {
		c.Ctx.WriteString("删除成功!")
	}
}

//ajax传递新用户数量
func (c *AdminController) Message() {
	id := c.Input().Get("id")
	num, _ := strconv.Atoi(id)
	message := models.GetNewUsers(num)
	newus := strconv.Itoa(message)
	c.Ctx.WriteString(newus)
}
