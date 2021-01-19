package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-module/carbon"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gbf")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(10)
	db.SetMaxIdleConns(6)
	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil
	}
	fmt.Println("connect success")
	return db
}

func GetDateTime() string {
	now := carbon.Now()
	m := now.Minute()
	switch {
	case m >= 40:
		m = 40
	case m >= 20:
		m = 20
	default:
		m = 0
	}
	return now.SetMinute(m).SetSecond(0).ToDateTimeString()
}
