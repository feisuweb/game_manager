#! /bin/bash
#默认进入的是登录用户的目录
cd /data/service/codeshop/codeshop_weixin/
tar -xzvf codeshop_weixin.tar.gz
#remove conf of dev
rm -rf /data/service/codeshop/codeshop_weixin/codeshop_weixin.tar.gz
nohup /data/service/codeshop/codeshop_weixin/codeshop_weixin &