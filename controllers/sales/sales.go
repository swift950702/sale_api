package sales

import (
	"api/conf"
	"api/controllers/account"
	model "api/models"
	"api/models/salestypes"
	"api/utils"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SalesController struct {
	beego.Controller
}

var db = model.DB

func (this *SalesController) AllSalesTypes() {
	definition, err := config.NewConfig("ini", "/root/lijun/sale_api/conf/app.conf")
	// definition, err := config.NewConfig("ini", "/Users/zuiyou/mywork/src/api/conf/app.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	fmt.Println(this.Ctx.Input)
	salestype := this.Ctx.Input.Param(":type")
	salestypeint, _ := strconv.Atoi(salestype)
	homeurl := definition.String("homeurl")
	u := &salestypes.SalesTypes{}
	saleslist := u.AllTypes(db, salestypeint)
	for _, i := range saleslist {
		temp := (*i).Cover
		(*i).Cover = homeurl + temp
	}
	Ret := utils.ReturnSuccess("", saleslist)

	this.Data["json"] = Ret
	this.ServeJSON()
}
func (this *SalesController) AddSalesTypes() {
	u := &salestypes.SalesTypes{}
	ret := u.InsertSale(db)
	this.Data["json"] = ret
	this.ServeJSON()
}
func (this *SalesController) SalesTypesById() {
	definition, err := config.NewConfig("ini", "/Users/zuiyou/mywork/src/api/conf/app.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	i := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(i)
	homeurl := definition.String("homeurl")
	u := &salestypes.SalesTypes{}
	sale := u.SalesById(db, id)
	temp := (*sale).Cover
	(*sale).Cover = homeurl + temp
	Ret := &account.Rps{}
	Ret.Data = *sale
	Ret.Errcode = conf.SUCCESS_CODE
	this.Data["json"] = Ret
	this.ServeJSON()
}
