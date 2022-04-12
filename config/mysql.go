package config

import (
	"fmt"
	"qp/back/helper"

	_ "github.com/go-sql-driver/mysql"
	gorm_sql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	sql_host := helper.GetEnv("MYSQL_HOST", "error")
	sql_port := helper.GetEnv("MYSQL_PORT", "3306")
	sql_user := helper.GetEnv("MYSQL_USER", "error")
	sql_pass := helper.GetEnv("MYSQL_PASSWORD", "error")
	sql_db := helper.GetEnv("MYSQL_DB", "error")

	// Format for the SQL CONNECTION: DNS
	// <user>:<password>@tcp(<host>:<port>)/<dbname>
	sqlConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", sql_user, sql_pass, sql_host, sql_port, sql_db)
	db, err := gorm.Open(gorm_sql.Open(sqlConnection), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
