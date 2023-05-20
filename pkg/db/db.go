package db

import (
	"apiserver/pkg/config"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(config *config.Config, logger *zap.Logger) (*gorm.DB, error) {
	dbConfig := config.DBConfig

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		logger.Error("connect to database error", zap.Error(err))
		return nil, err
	}
	sqlDb, err := db.DB()

	if err != nil {
		return nil, err
	}
	sqlDb.SetConnMaxLifetime(time.Hour)
	sqlDb.SetMaxOpenConns(dbConfig.MaximumIdleSize)
	sqlDb.SetMaxIdleConns(dbConfig.MaximumIdleSize)

	return db, nil
}
