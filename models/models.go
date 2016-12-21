package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"math"
)

type User struct {
	Id       int64
	Name     string
	Account  string
	PhoneNum string
	Created  string `orm:"index"`
}

var DefaultRowsLimit = 1000

//注册数据库
func RegisterDB() {
	//注册模型
	orm.RegisterModel(new(User))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册数据库
	orm.RegisterDataBase("default", "mysql", "root:special005@/user?charset=utf8", 10)
}

//获取新用户
func GetNewUsers(id int) int {
	o := orm.NewOrm()
	users := make([]*User, 0)
	qs := o.QueryTable("user")
	qs.Limit(10, id).All(&users)
	return len(users)
}

//获取所有用户资料
func GetAllUsers() ([]*User, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	qs := o.QueryTable("user")
	_, err := qs.All(&users)
	return users, err
}

//删除用户
func DeleteUser(id int64) error {
	var err error
	o := orm.NewOrm()
	user := &User{Id: id}
	_, err = o.Delete(user)
	if err != nil {
		return err
	}
	return err
}

//获取数据
func GetUsers(page, prepage int) (map[string]interface{}, error) {
	//获取数据库信息
	o := orm.NewOrm()
	allusers := make([]*User, 0)
	qs := o.QueryTable("user")
	qs.All(&allusers)
	/********************分页************************/
	nums := len(allusers)
	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int //所有页
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
	}
	//获取数据
	users := make([]*User, 0)
	startP := (page - 1) * prepage
	_, err := qs.Limit(prepage, startP).OrderBy("-created").All(&users)
	if err != nil {
		return nil, err
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["users"] = users
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	paginatorMap["totals"] = nums
	return paginatorMap, err
}
