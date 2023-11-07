package dbmysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config mysql配置
type Config struct {
	User     string // 用户
	Password string // 密码
	Host     string // 地址
	Port     int    // 端口
	Database string // 数据库
}

// NewDB 创建mysql连接
func NewDB(c *Config) (*gorm.DB, error) {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

// MustNewDB 创建mysql连接
func MustNewDB(c *Config) *gorm.DB {
	db, err := NewDB(c)
	if err != nil {
		panic(fmt.Errorf("dbmysql connect err:%v", err))
	}
	return db
}
