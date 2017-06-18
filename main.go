package main

import (
	"github.com/astaxie/beego"
	_ "github.com/feisuweb/game_manager/models"
	_ "github.com/feisuweb/game_manager/routers"
	"os"
)

func main() {
	//创建附件目录
	os.Mkdir("logs", os.ModePerm)
	os.Mkdir("upload", os.ModePerm)
	os.Mkdir("upload/images", os.ModePerm)
	os.Mkdir("upload/files", os.ModePerm)
	beego.SetLogFuncCall(true)
	beego.SetLogger("file", `{"filename":"logs/web.log"}`)
	beego.Info("服务已经启动...")
	beego.Run()
}
