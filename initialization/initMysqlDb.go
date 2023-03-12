package initialization

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var mysqlDb *sql.DB

func InitMysql() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"))
	mysqlDb, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//_, err = open.Exec("CREATE DATABASE sql_test ;")
	//if err != nil {
	//	zap.L().Error("创建数据库错误: " + err.Error())
	//	return
	//}
	//_, err = open.Exec("USE sql_test ;")
	//if err != nil {
	//	zap.L().Error("使用数据库错误: " + err.Error())
	//	return
	//}
	return nil
}
