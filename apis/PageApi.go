package apis

import (
	"github.com/feisuweb/game_manager/models"
)

//资讯列表模型
type ApiPageList struct {
	PageList    []*models.Page //资讯列表
	PageNumber  int64
	RecordCount int64
}

//资讯详细页模型
type ApiPageDetail struct {
	Detail *models.Page //资讯详情
}
