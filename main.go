package main

import (
	"github.com/aa8616436/goblog/pkg/setting"
	"github.com/aa8616436/goblog/routers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	//初始化配置
	setting.InitSetting()

	//初始化数据库
	go setting.InitDatabase()
	//初始化缓存
	go setting.InitRedisPool()
	//初始化Http gin
	go routers.InitHttp()

	//起到阻塞的作用
	select {}
}
