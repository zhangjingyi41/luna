package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name      string `mapstructure:"name"`       // 应用名称
	Mode      string `mapstructure:"mode"`       // 运行模式
	Port      int    `mapstructure:"port"`       // 端口号
	Version   string `mapstructure:"version"`    // 版本号
	StartTime string `mapstructure:"start_time"` // 启动时间
	MachineId int64  `mapstructure:"machine_id"` // 机器ID

	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*AuthConfig  `mapstructure:"auth"`
}

type AuthConfig struct {
	JwtExpire int `mapstructure:"jwt_expire"` // JWT过期时间
}

type LogConfig struct {
	Level      string `mapstructure:"level"`       // 日志级别
	Filename   string `mapstructure:"filename"`    // 日志文件名
	MaxSize    int    `mapstructure:"max_size"`    // 日志文件最大大小
	MaxAge     int    `mapstructure:"max_age"`     // 日志文件最大保存时间
	MaxBackups int    `mapstructure:"max_backups"` // 日志文件最大保存数量
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`           // 数据库地址
	Port         int    `mapstructure:"port"`           // 数据库端口
	User         string `mapstructure:"user"`           // 数据库用户名
	Password     string `mapstructure:"password"`       // 数据库密码
	Dbname       string `mapstructure:"dbname"`         // 数据库名称
	MaxOpenConns int    `mapstructure:"max_open_conns"` // 最大连接数
	MaxIdleConns int    `mapstructure:"max_idle_conns"` // 最大空闲连接数
	Prefix       string `mapstructure:"prefix"`         // 表前缀
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`      // 数据库地址
	Port     int    `mapstructure:"port"`      // 数据库端口
	Password string `mapstructure:"password"`  // 数据库密码
	Db       int    `mapstructure:"db"`        // 数据库编号
	PoolSize int    `mapstructure:"pool_size"` // 连接池大小
}

var App = new(AppConfig)

func InitConfig(filePath string) (err error) {
	fmt.Println("配置文件加载中......")
	viper.SetConfigFile(filePath)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	if err := viper.Unmarshal(App); err != nil {
		fmt.Printf("读取配置文件失败: %v\n", err)
	}

	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("监听到配置文件修改\n")
		if err := viper.Unmarshal(App); err != nil {
			fmt.Printf("配置文件解析失败: %v\n", err)
		}
	})
	fmt.Println("配置文件加载完成")
	return nil
}
