package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"student/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// 第 1 步引入 *gorm.DB
type Data struct {
	// TODO wrapped database client
	gormDB *gorm.DB
}

// 第 2 步初始化 gorm
func NewGormDB(c *conf.Data) (*gorm.DB, error) {
	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(150)
	sqlDB.SetConnMaxLifetime(time.Second * 25)
	return db, err
}

// 第 3 步，初始化 Data
func NewData(logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{gormDB: db}, cleanup, nil
}

// 第 4 步，用 wire 注入代码，修改 原来的 NewSet
var ProviderSet = wire.NewSet(NewData, NewGormDB, NewStudentRepo)
