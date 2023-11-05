package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	User     string
	PassWord string
	DB       string
}

func NewMysql(c *Config) (*gorm.DB, error) {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.PassWord, c.Host, c.Port, c.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func MustNewMysql(c *Config) *gorm.DB {
	db, err := NewMysql(c)
	if err != nil {
		panic(fmt.Errorf("mysql connect err:%v", err))
	}
	return db
}
