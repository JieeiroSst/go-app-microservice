package mysql

import (
	"fmt"
	"github.com/JIeeiroSst/go-app/config"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConn(c *config.Config) (*gorm.DB,error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			c.Mysql.MysqlUser,
			c.Mysql.MysqlPassword,
			c.Mysql.MysqlHost,
			c.Mysql.MysqlPort,
			c.Mysql.MysqlDbname,
		)
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "gorm.Connect")
	}
	err = db.AutoMigrate()
	if err != nil {
		return nil, errors.Wrap(err, "gorm.Connect")
	}
	return db,nil
}
