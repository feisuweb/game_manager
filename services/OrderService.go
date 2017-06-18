package services

import (
	"fmt"
	"github.com/feisuweb/game_manager/models"
	"math/rand"
	"strings"
	"time"
)

type OrderService struct {
}

//生成订单号
func (m *OrderService) GetRandOrderNo() string {

	datetime := time.Now().Format("2006-01-02 15:04:05")
	datetime = strings.Replace(datetime, " ", "", -1)
	datetime = strings.Replace(datetime, ":", "", -1)
	datetime = strings.Replace(datetime, "-", "", -1)

	var l int64
	l = 8
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	rn := string(result)

	no := fmt.Sprintf("%s%sP", datetime, rn)
	return no
}

func (m *OrderService) FindOrderByProductIdWithMemberId(productId int64, memberId int64) bool {

}

func (m *OrderService) FindOrderByOrderNo(orderNo string) bool {

}
func (m *OrderService) CreateOrder(orderInfo Order) (int64, bool) {

	var (
		info        *models.Order   = new(models.Order)
		productInfo *models.Product = new(models.Product)
		minfo       *models.Member  = new(models.Member)
		payinfo     *models.PayLog  = new(models.PayLog)
		agentInfo   *models.Agent   = new(models.Agent)
		err         bool
		orderNo     string
		url         string
		member_id   int64
		product_id  int64
		mobile      string
		email       string
		ip          string
	)

	member_id, _ = this.GetInt64("member_id")
	product_id, _ = this.GetInt64("product_id")
	mobile = strings.TrimSpace(this.GetString("mobile"))
	email = strings.TrimSpace(this.GetString("email"))

	if !utils.ValidateMobile(mobile) {
		this.Abort("手机号格式错误！")
		return
	}
	if !utils.ValidateEmail(email) {
		this.Abort("请填写正确格式的邮箱！")
		return
	}
	ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	//检查用户之前是否注册过本网站，注册过，则直接登录
	err = minfo.FindMemberByMobileAndEmail(mobile, email)
	if err {
		//如果查询到用户已经存在，则
		member_id = minfo.Id
	} else {
		//注册账号信息
		//默认以邮箱和手机号注册一个用户，用户密码是随机数
		//username string, password string, mobile string, email string, ip string
		//ipResult := models.TabaoAPI(ip)
		password := "123456"
		minfo.Email = email
		minfo.Password = password
		minfo.Mobile = mobile
		minfo.Nickname = "会员" + mobile
		minfo.MemberName = mobile
		minfo.RegisterIp = ip
		minfo.IsVip = 0
		minfo.IsValidateMobile = 0
		minfo.IsValidateEmail = 0
		minfo.Points = 0
		minfo.Money = 0
		err := minfo.Register()
		if err {
			member_id = minfo.Id
		}
	}
	//根据产品ID查询产品信息

	err = productInfo.FindProductById(product_id)

	if !err {
		this.Abort("产品信息有误，请查验后再提交")
	}

	if productInfo.Price == 0 {
		this.Abort("此产品是免费下载，无需购买！")
	}

	//如果是VIP会员，则直接判断
	if minfo.CheckVip(minfo.Id) {
		//VIP 会员，采用会员价购买
		if productInfo.VipPrice == 0 {
			this.Abort("此产品是会员免费下载，无需购买！")
		}
	}
	//判断之前是否已经购买过，购买过则无需再次购买
	err = info.FindOrderByProductIdWithMemberId(product_id, member_id)

	if err {
		if info.Status > 0 {
			mid2 := fmt.Sprintf("%d", member_id)
			this.Ctx.SetCookie("member_id", mid2)
			url := fmt.Sprintf("/download/%d.html?orderno=%s", info.ProductId, info.OrderNo)
			this.Redirect(url, 302)
			return
		}
	}

	//订单创建流程开始
	//获取随机订单号
	orderNo = info.GetRandOrderNo()
	//订单创建

	info.OrderNo = orderNo
	info.ProductId = product_id
	info.ProductName = productInfo.Name
	info.MemberId = member_id
	info.FromPlatform = "pc"
	info.FromChannel = "direct"
	info.FromChannelTag = "codeshop.com"

	info.RecommendCode = agentInfo.RecommendCode
	info.AgentId = agentInfo.Id
	info.AgentName = agentInfo.AgentName
	info.AgentWeixinOpenId = agentInfo.WeixinOpenId
	info.AgentWeixin = agentInfo.Weixin
	info.AgentEmail = agentInfo.Email
	info.AgentMobile = agentInfo.Mobile

	info.MemberName = minfo.Nickname
	info.MemberMobile = minfo.Mobile
	info.MemberEmail = minfo.Email
	info.MemberWeixin = minfo.Weixin
	info.MemberWeixinOpenId = minfo.WeixinOpenId

	info.MemberName = minfo.Nickname
	info.MemberMobile = mobile
	info.MemberEmail = email
	info.MemberWeixin = minfo.Weixin
	info.CommissionAmount = 0

	if minfo.CheckVip(member_id) {
		//VIP 会员，采用会员价购买
		info.Price = productInfo.VipPrice
		info.Discount = productInfo.Price - productInfo.VipPrice
	} else {
		//普通会员，采用普通价格购买
		info.Price = productInfo.Price
		info.Discount = 0
	}

	info.PayAmount = info.Price
	info.Amount = info.Price
	info.IsSend = 0
	info.Status = 0
	//创建订单

	orderId, oerr := info.CreateOrder()

	if oerr {
		//创建微信支付记录
		payinfo.OrderId = orderId
		payinfo.OrderNo = info.OrderNo
		payinfo.PayType = 1 //消费
		payinfo.MemberId = member_id
		payinfo.AgentId = info.AgentId
		payinfo.MemberName = info.MemberName
		payinfo.MemberMobile = mobile
		payinfo.MemberEmail = email
		payinfo.MemberWeixin = info.MemberWeixin
		payinfo.Amount = info.Amount
		payinfo.Discount = info.Discount
		payinfo.PayAmount = info.PayAmount
		payinfo.PayMethod = "weixin"
		payinfo.PayBody = "购买" + info.ProductName + "-优品源码网"
		payinfo.ProductId = info.ProductId
		payinfo.PayStatus = 0
		payinfo.Status = 0
		payinfo.Insert()
	}

	url = site_pay_scan_url + "?orderno=" + orderNo
	if info.PayAmount > 0 {
		url = site_pay_scan_url + "?orderno=" + orderNo
	} else {
		//直接跳转会产品页面
		url = fmt.Sprintf("/product/%d.html", product_id)
	}
	//页面cache控制
	this.Ctx.Output.Header("Cache-Control", "public")
	CategoryList := models.GetCategoryList()
	this.Data["CategoryList"] = CategoryList
	mid3 := fmt.Sprintf("%d", member_id)
	this.Ctx.SetCookie("member_id", mid3)
	this.Ctx.SetCookie("token", minfo.Password)
	this.Redirect(url, 302)
	return 0, false
}
