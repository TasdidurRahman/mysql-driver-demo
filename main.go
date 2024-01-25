package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	cfg, err := mysql.ParseDSN("tcp(mysql-demo.mysql.svc:3306)/")
	if err != nil {
		panic(err)
	}
	cfg.User = "user"
	cfg.Passwd = "password"
	// cfg.DBName = "adb"
	fmt.Println(cfg.FormatDSN())
	//
	//_, err := sql.Open("mysql", "tcp(mysql-demo.mysql.svc:3306)/")
	//if err != nil {
	//	panic(err)
	//}

	ParseAll()
}

/*
output

user:password@tcp(mysql-demo.mysql.svc:3306)/
====================================================
testing for url :  tcp(mysql-demo.mysql.svc:3306)/
error :  parse "tcp(mysql-demo.mysql.svc:3306)/": first path segment in URL cannot contain colon
====================================================
testing for url :  tcp://mysql-demo.mysql.svc:3306
url: tcp://mysql-demo.mysql.svc:3306
host: mysql-demo.mysql.svc:3306
hostname: mysql-demo.mysql.svc
port: 3306
scheme: tcp
rebuild mysql connection from url2 :  uname:pword@tcp(mysql-demo.mysql.svc:3306)/
====================================================
*/
