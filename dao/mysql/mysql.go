package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sql.DB

func Init() (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		zap.L().Error("Cannot connect to mysql", zap.Error(err))
		return err
	}

	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return err
}

func Close() {
	db.Close()
}
