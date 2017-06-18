package controllers

type HttpErrorHandle struct {
	baseController
}

func (this *HttpErrorHandle) Error404() {
	this.TplName = "_error.html"
}

func (this *HttpErrorHandle) Error501() {

	this.TplName = "_error.html"
}

func (this *HttpErrorHandle) ErrorDb() {
	this.TplName = "_error.html"
}
