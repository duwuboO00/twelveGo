package database

import (
	"database/sql"
	"fmt"

	"twelveGo/config"

	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

// DBURL 為全域變數，存放組合好的連線字串
var DBURL string

// DB 為連線池 singleton，供其他模組共用
var DB *sql.DB

func InitDB() (*sql.DB, error) {
	cfg := config.GetConfig() // 從 config.go 取得設定
	pg := cfg.PostgresConfig
	DBURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pg.Host, pg.Port, pg.User, pg.Password, pg.DBName, pg.SSLMode)

	db, err := sql.Open("postgres", DBURL)
	if err != nil {
		zap.L().Fatal("無法連接到數據庫", zap.Error(err))
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(0)

	if err := db.Ping(); err != nil {
		zap.L().Fatal("無法 ping 通數據庫", zap.Error(err))
		return nil, err
	}

	DB = db

	zap.L().Info("數據庫連接成功！")
	return db, nil
}
