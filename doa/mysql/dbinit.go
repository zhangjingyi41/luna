package mysql

import (
	"fmt"
	"luna/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init(c *conf.MysqlConfig) (err error) {
	fmt.Println("数据库配置加载中......")
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Dbname)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Prefix,
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	fmt.Println("数据库配置加载完成")
	return
}

func Close() {
	d, _ := db.DB()
	d.Close()
}
