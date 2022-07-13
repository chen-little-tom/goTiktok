/**
 * @Author: chenchen
 * @Description:
 * @File: initDao
 * @Version: 1.0.0
 * @Data: 2022/07/13 21:49
 */

package dao

import (
	"douyin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func Init() {
	newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      true,
		})
	var err error
	dsn := config.MysqlURL
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newlogger,
	})
	if err != nil {
		log.Panicln("err: ", err.Error())
	}
}
