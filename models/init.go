package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/feisuweb/game_manager/libs/utils"
	"github.com/go-xorm/xorm"
	_ "github.com/lunny/godbc"
	"path/filepath"
	//"time"
)

var engine *xorm.Engine

func init() {
	configPath := ""
	fmt.Println("current run mode is " + beego.AppConfig.String("runmode"))
	//读取配置文件
	if beego.AppConfig.String("runmode") == "dev" {
		configPath = filepath.Join("conf", "dev.database.conf")
	} else {
		configPath = filepath.Join("conf", "pro.database.conf")
	}

	fmt.Println("Config path:" + configPath)
	red, err := utils.GetConfig(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	//读取database 配置
	databasepass := red.Conf["database.password"]
	databaseuser := red.Conf["database.user"]
	databasedb := red.Conf["database.database"]
	databasehost := red.Conf["database.host"]
	databaseport := red.Conf["database.port"]
	//密码长度，判断是否已经加密过
	if len(databasepass) == 24 {
		databasepass, err = utils.Decrypt(databasepass)
		if err != nil {
			fmt.Errorf("Decrypt database passwd failed.")
			return
		}
	}
	//没有加密密码，则加密一次密码，并写入配置文件
	if len(databasepass) != 24 {
		psd, err := utils.Encrypt(databasepass)
		if err != nil {
			fmt.Errorf("decrypt passwd failed.%v", psd)
			return
		}
		psd = "\"" + psd + "\""
		red.Set("database.password", psd)
	}

	engine, err := xorm.NewEngine("odbc", "driver={SQL Server};Server="+databasehost+";Database="+databasedb+";port="+databaseport+"; uid="+databaseuser+"; pwd="+databasepass+";")
	engine.SetMaxIdleConns(5)
	err = engine.Sync(new(VipMachine))

}
