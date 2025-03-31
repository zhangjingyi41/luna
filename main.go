package main

import (
	"fmt"
	"luna/conf"
	"luna/doa/mysql"
	"luna/logger"
	"os"
)

func main() {
	// 加载配置文件
	var filepath string
	if len(os.Args) < 2 {
		// 如果参数小于2，则默认使用luna.yml
		filepath = "config/luna-dev.yml"
	} else {
		filepath = os.Args[1]
	}
	if err := conf.InitConfig(filepath); err != nil {
		fmt.Printf("初始化配置失败: %v\n", err)
		return
	}

	// 初始化日志配置
	if err := logger.Init(conf.App.LogConfig, conf.App.Mode); err != nil {
		fmt.Printf("初始化日志失败: %v\n", err)
		return
	}

	// 初始化数据库
	if err := mysql.Init(conf.App.MysqlConfig); err != nil {
		fmt.Printf("初始化数据库失败: %v\n", err)
		return
	}
	defer mysql.Close() // 程序关闭时断开数据库链接

}
