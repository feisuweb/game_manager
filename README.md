codeshop 是源码交易网站,采用golang编写，使用beego 1.8开发，golang 1.8编译 

#环境
golang 1.8+ 

# 数据库

数据采用mysql数据库,自动生成表结构

# 数据安全

* 数据库配置自动aes加密
* 用户密码采用多层md5+随机串加密
* 短信登录后台控制.

# 模块说明
1. codeshop_weixin 管理后台,独立应用  admin.codeshop.com/ypadmin   端口 8010
- codeshop_web 网站前端,独立应用 www.codeshop.com    端口 8020
- codeshop_api api接口,json格式,独立应用 api.codeshop.com  端口  8050
- codeshop_mobile 手机网页,独立应用 m.codeshop.com   端口 8040
- codeshop_pay  微信扫码支付应用,独立应用 pay.codeshop.com 端口 8030
- game_manager 公共模块,无法直接编译. 
- codeshop_task 定时任务模块,主要是实现订单发货,短信通知等内容. 无外网访问域名 端口 8070
- codeshop_weixin 微信公众号模块,用来处理公众号的请求,独立应用 weixin.codeshop.com   8060

