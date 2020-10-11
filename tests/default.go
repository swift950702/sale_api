package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	Name  string
	Phone string
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	con := session.DB("test").C("people")
	err = con.Insert(&Person{"Ale", "111111"}, &Person{"Cla", "222222222"})
	if err != nil {
		panic(err)
	}
	result := Person{}
	err = con.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone:", result.Phone)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) GetRedis() {
	coon, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return
	}
	rs, err := coon.Do("SET", "lijun", "hello")
	fmt.Println(rs)
	defer coon.Close()
}
