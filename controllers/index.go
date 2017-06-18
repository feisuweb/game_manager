package controllers

import (
//"encoding/json"

//"strconv"
//"time"
//"github.com/feisuweb/game_manager/filters"
//"github.com/feisuweb/game_manager/models"
)

// 前台页面handle
type IndexHandle struct {
	baseController
}

///前台首页
func (this *IndexHandle) Index() {
	this.Layout = "layout/_default_layout.html"
	this.TplName = "_index.html"
}
