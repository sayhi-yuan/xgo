package core

import (
	"xgo/config"
	"context"
	"fmt"
	"time"

	"git.qdreads.com/gotools/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func init() {
	initMysql()
}

var db *gorm.DB

func initMysql() {
	// 获取mysql配置
	cfg := config.Cfg.MySQL.Default

	// 构建dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DataBase, cfg.Config)

	// 初始化gorm
	if orm, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 关闭复数表
		},
		// 注入日志服务
		Logger: &log.GormLogger{
			LogLevel: logger.Info,
		},
	}); err != nil {
		panic(1)
	} else {
		// 设置gorm链接池
		if d, err := orm.DB(); err != nil {
			panic("配置gorm链接池失败")
		} else {
			d.SetMaxIdleConns(cfg.MaxIdleConn)
			d.SetMaxOpenConns(cfg.MaxOpenConn)
			d.SetConnMaxLifetime(time.Hour)
		}

		// 是否开启debug模式
		if cfg.Debug {
			db = orm.Debug()
		} else {
			db = orm
		}
	}
}

func DB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
