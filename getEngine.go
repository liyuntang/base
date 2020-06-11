package base

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_"github.com/go-sql-driver/mysql"
)

func GetEngine(endPoint, database string) *xorm.Engine {
	// 获取数据库连接
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true", "admin_m", "XphoeGz6NED4yQMQ0yJc", endPoint, database, "utf8mb4")
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		return nil
	}
	return engine
}
