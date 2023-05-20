package gorms

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)
var Msyqllogger logger.Interface

func init() {
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	Dbname := "inventory"
	timeout := "10s"
	Msyqllogger = logger.Default.LogMode(logger.Info)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username,
		password, host, port, Dbname, timeout)
	//连接mysql，获得DB类型实例，用于读写操作
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ //默认跳过事务
		//SkipDefaultTransaction: true,
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix:   "g_", //表名前缀
		//	SingularTable: true, //是否单数单数表名
		//	NoLowerCase:   true, //是：关闭大小写转换 否：进行大小写转换 默认false
		//},
		Logger: Msyqllogger,
	})
	if err != nil {
		panic("数据库连接失败，err=" + err.Error())
	}
	//连接成功
	DB = db

}
