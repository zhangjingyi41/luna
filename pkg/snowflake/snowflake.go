package snowflake

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

// 初始化雪花算法
func Init(startTime string, machineId int64) (err error) {
	fmt.Println("雪花算法初始化中......")
	var tmp time.Time
	// 解析开始时间
	if tmp, err = time.Parse("2006-01-02", startTime); err != nil {
		return err
	}
	// 设置开始时间
	snowflake.Epoch = tmp.UnixNano() / 1000000
	// 创建雪花算法节点
	if node, err = snowflake.NewNode(machineId); err != nil {
		return err
	}
	fmt.Println("雪花算法初始化完成")
	return nil
}

// 生成ID
func GenID() int64 {
	return node.Generate().Int64()
}
