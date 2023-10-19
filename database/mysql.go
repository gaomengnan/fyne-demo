package database

import (
	"database/sql"

	"github.com/gaomengnan/fyne-demo/data"
	// 导入 MySQL 驱动程序
	_ "github.com/go-sql-driver/mysql"
)

func init() {

}

func TestConnect(conf *data.SerializationConnectionData) error {
	// 数据库连接信息
	db, err := sql.Open("mysql", conf.DSN())

	if err != nil {
		return err
	}
	// 一定要在函数结束时关闭数据库连接
	defer db.Close()

	// 测试数据库连接是否正常
	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}
func Connect(conf *data.SerializationConnectionData) error {
	// 数据库连接信息
	db, err := sql.Open("mysql", conf.DSN())

	if err != nil {
		return err
	}
	// 一定要在函数结束时关闭数据库连接
	defer db.Close()

	// 测试数据库连接是否正常
	err = db.Ping()
	if err != nil {
		return err
	}

	conf.Connection = db
	return nil
}
