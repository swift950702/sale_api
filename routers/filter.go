package routers

import (
	"encoding/json"

	"api/xlib/auth"

	"github.com/astaxie/beego/context"
	"github.com/gbrlsnchs/jwt/v3"
)

type TokenPayload struct {
	jwt.Payload
	Pwd      string `json:"pwd"`
	Username string `json:"username"`
}

var hs = jwt.NewHS256([]byte("secret"))
var filterLogin = func(ctx *context.Context) {
	token := ctx.Input.Header("Authorization")
	_, err := auth.Verify([]byte(token))
	if err != nil {
		jsonData := make(map[string]interface{})
		jsonData["errcode"] = 403
		jsonData["data"] = "请登录后再操作"
		returnJSON, _ := json.Marshal(jsonData)
		ctx.ResponseWriter.Write(returnJSON)
	}
}

func init() {
	// beego.InsertFilter("/v1/*", beego.BeforeRouter, filterLogin)
}
