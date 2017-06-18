package routers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/feisuweb/game_manager/controllers"
	//"github.com/feisuweb/game_manager/filters"
)

func init() {
	//pages
	beego.Router("/", &controllers.IndexHandle{}, "*:Index")
	///error handel
	beego.ErrorController(&controllers.HttpErrorHandle{})
}
