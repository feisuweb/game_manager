#! /bin/bash
echo 'build for linux'
bee pack -be GOOS=linux -exp=.:upload:logs:bee.json:README.md:shell  
echo 'upload to linux server'
scp codeshop_weixin.tar.gz root@www.wqdsoft.com:/data/service/codeshop/codeshop_weixin/
ssh root@www.wqdsoft.com 'bash -s' < ./shell/deploy_restart.sh