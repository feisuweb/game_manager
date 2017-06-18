package controllers

import (
	"github.com/astaxie/beego"
)

func init() {

}

type baseController struct {
	beego.Controller
}

func (this *baseController) Prepare() {

}

func Error(err error) {
	if err != nil {
		panic(err)
		beego.Error(err.Error())
		//os.Exit(1)
	}
}
