package db

import (
	"elotus/pkg/cfg"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	MySQLURI      = "MYSQL_URI"
	MySQLDatabase = "MYSQL_DATABASE"
	MySQLUsername = "MYSQL_USERNAME"
	MySQLPassword = "MYSQL_PASSWORD"

	MySQLMaxConnection = "MYSQL_MAX_CONNECTION"
)

func NewMySQLClient() *gorm.DB {
	connStr := mysqlDatabaseConfiguration()

	var engine *gorm.DB
	var err error

	if cfg.Reader().MustGetString("SERVER_MODE") == "debug" {
		engine, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})
	} else {
		engine, err = gorm.Open(mysql.Open(connStr), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             time.Second,
					LogLevel:                  logger.Silent,
					IgnoreRecordNotFoundError: true,
					ParameterizedQueries:      true,
					Colorful:                  false,
				}),
			SkipDefaultTransaction: true,
		})
	}

	if err != nil {
		panic(fmt.Sprintf("could not initialize mysql client: %s", err.Error()))
	}

	db, _ := engine.DB()

	max := cfg.Reader().MustGetInt(MySQLMaxConnection)
	if max == 0 {
		max = 8
	}

	db.SetMaxOpenConns(max)
	db.SetConnMaxIdleTime(2 * time.Minute)
	db.SetConnMaxLifetime(time.Hour)

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("could not ping mysql db: %s", err.Error()))
	}

	return engine
}

func mysqlDatabaseConfiguration() string {
	uri := cfg.Reader().MustGetString(MySQLURI)
	database := cfg.Reader().MustGetString(MySQLDatabase)
	username := cfg.Reader().MustGetString(MySQLUsername)
	password := cfg.Reader().MustGetString(MySQLPassword)

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		username, password, uri, database)
}
