package apis

import (
	"github.com/feisuweb/game_manager/models"
)

///详情页输出结构
type ApiDetail struct {
	ProductDetail     *models.Product           //产品信息
	ProductImageList  []*models.ProductImage    //图片列表
	ProductSourceCode *models.ProductSourceCode //源码信息
	Seller            *models.Seller            //卖家信息
	Supplier          *models.Supplier          //供应商信息
	ProductList       []*models.Product         //类似产品
	ProductClass      *models.ProductClass      //分类信息
}

//今日更新列表
type ApiToday struct {
	ProductList []*models.Product //产品列表
}

//列表页输出模型
type ApiList struct {
	ProductClass *models.ProductClass //分类信息
	ProductList  []*models.Product    //产品列表
	PageNumber   int64                //当前页码
	RecordCount  int64                //记当总数
}

type ApiIndexProductList struct {
	ProductList []*models.Product //产品各种小列表
}

//导航菜单接口
type ApiNavigation struct {
	NavigationList []*models.ProductClass //导航列表
}

//手机导航菜单接口
type ApiMenu struct {
	MenuList []*models.ProductClass //导航列表
}

//首页输出模型
type ApiIndex struct {
	List []*ApiIndexProductList
}

//搜索模型
type ApiSearch struct {
	ProductList []*models.Product //产品列表
	PageNumber  int64
	RecordCount int64
}

//产品推荐模型
type ApiNew struct {
	ProductList []*models.Product //产品列表
}
