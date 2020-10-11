package routers

import (
	"api/controllers/account"
	"api/controllers/sales"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/testmgo", &controllers.MainController{})
	// beego.Router("/testredis", &controllers.MainController{}, "*:GetRedis")
	beego.Router("/v1/user/info", &account.AccountController{}, "get:GetInfo")
	beego.Router("/v1/salestype/detail/:type", &sales.SalesController{}, "post:AllSalesTypes")
	beego.Router("/v1/salestype/id/:id", &sales.SalesController{}, "post:SalesTypesById")
	beego.Router("/v1/salestype/addsale", &sales.SalesController{}, "post:AddSalesTypes")

	//第一个参数url路径，第二个参数是应用下文件夹的名称，beego默认将static设为静态文件夹，不用配置
	// 注册静态文件目录
	beego.SetStaticPath("/images/index", "static/img")

	beego.Router("/user/login", &account.AccountController{}, "*:Login")
}
