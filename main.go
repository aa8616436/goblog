package main

import (
	"github.com/aa8616436/goblog/pkg/setting"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	//初始化配置
	setting.InitSetting()
	//
}
