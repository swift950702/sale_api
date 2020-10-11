package mylogs

import (
	"github.com/astaxie/beego/logs"
)

func Info(str string) {
	logs.SetLogger(logs.AdapterFile, `{"filename":"mylogs/info.log"}`)
	logs.Info(str)

}
func error(str string) {
	logs.SetLogger(logs.AdapterFile, `{"filename":"mylogs/error.log"}`)
	logs.Error(str)
}
