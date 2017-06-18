package filters

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	//"github.com/feisuweb/game_manager/models"
)

// func IsLogin(ctx *context.Context) (bool, models.Member) {
// 	token, flag := ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
// 	var memberInfo models.Member
// 	if flag {
// 		flag, memberInfo = models.FindMemberByToken(token)
// 	}
// 	return flag, memberInfo
// }

var FilterUser = func(ctx *context.Context) {
	//ok, _ := IsLogin(ctx)
	ok := true
	if !ok {
		ctx.Redirect(302, "/member/login")
	}
}
