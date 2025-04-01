package main

import (
	"fmt"
	"luna/conf"
	"luna/doa/mysql"

	docs "luna/docs"
	"luna/logger"
	"luna/pkg/snowflake"
	"luna/router"
	"os"
)

// @title luna
// @version 0.0.1
// @description markdown博客系统
// @host localhost:8084
// @license.name MIT
// @license.url https://github.com/zhangjingyi41/luna/blob/master/LICENSE
// @BasePath /
func main() {
	// 加载配置文件
	var filepath string
	if len(os.Args) < 2 {
		// 如果参数小于2，则默认使用luna.yml
		filepath = "conf/luna-dev.yml"
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

	// 初始化雪花算法
	if err := snowflake.Init(conf.App.StartTime, conf.App.MachineId); err != nil {
		fmt.Printf("初始化雪花算法失败: %v\n", err)
		return
	}

	docs.SwaggerInfo.BasePath = "/"
	// 初始化http服务
	server := router.RunServer(conf.App.Mode)

	// 启动服务
	if err := server.Run(fmt.Sprintf(":%d", conf.App.Port)); err != nil {
		fmt.Printf("启动服务失败: %v\n", err)
		return
	}
}
