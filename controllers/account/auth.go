package account

import (
	"api/conf"
	"api/models/users"
	"api/mylogs"
	"api/xlib/auth"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AccountController struct {
	beego.Controller
}

// func (c *AccountController) Get() {
// 	session, err := mgo.Dial("")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	session.SetMode(mgo.Monotonic, true)
// 	con := session.DB("test").C("people")
// 	err = con.Insert(&Person{"Ale", "111111"}, &Person{"Cla", "222222222"})
// 	if err != nil {
// 		panic(err)
// 	}
// 	result := Person{}
// 	err = con.Find(bson.M{"name": "Ale"}).One(&result)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Phone:", result.Phone)
// }

// 根据token获取用户信息
func (c *AccountController) GetInfo() {

}

// 用户登录
func (c *AccountController) Login() {
	fmt.Println(c.Ctx.Input.Header("Authorization"))
	db, _ := gorm.Open("mysql", "root:lijun950702@/domesticDB?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	u := &users.User{}
	Ret := &Rps{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, u); err == nil {
		mylogs.Info(fmt.Sprintf("Login success by param：%v", *u))
		ok, code := u.SearchByName(db)
		if ok {
			retStr, _ := auth.Sign(u.Username, u.Pwd)
			Ret.Errcode = conf.SUCCESS_CODE
			Ret.Data = retStr
		} else if code == 2 {
			Ret.Errcode = conf.PWD_ERR
			Ret.Data = "密码错误"
		} else {
			Ret.Errcode = conf.DB_NOTFOUND
			Ret.Data = "没有这条数据"
		}
	}
	c.Data["json"] = Ret
	c.ServeJSON()
}
func (c *AccountController) GetRedis() {
	coon, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return
	}
	rs, err := coon.Do("SET", "lijun", "hello")
	fmt.Println(rs)
	defer coon.Close()
}
