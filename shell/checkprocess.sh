#! /bin/bash
# program : 判断进行是否存在，并重新启动
function check(){
  count=`ps -ef |grep $1 |grep -v "grep" |wc -l`
  #echo $count
  if [ 0 == $count ];then
    nohup /data/service/codeshop/$1/$1 &
  fi
}

check codeshop_weixin