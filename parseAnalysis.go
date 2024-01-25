package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"net/url"
)

func ParseAll() {

	fmt.Println("====================================================")
	url1 := "tcp(mysql-demo.mysql.svc:3306)/"
	fmt.Println("testing for url : ", url1)

	u1, err := url.Parse(url1)
	if err != nil {
		fmt.Println("error : ", err)
	} else {
		fmt.Println(u1)
	}

	fmt.Println("====================================================")
	url2 := "tcp://mysql-demo.mysql.svc:3306"
	fmt.Println("testing for url : ", url2)

	u2, err := url.Parse(url2)
	if err != nil {
		fmt.Println("error : ", err)
	} else {

		fmt.Printf("url: %s\nhost: %s\nhostname: %s\nport: %s\nscheme: %s\n", u2, u2.Host, u2.Hostname(), u2.Port(), u2.Scheme)
		fmt.Println("rebuild mysql connection from url2 : ", RebuildDSN("uname", "pword", url2))
	}
	fmt.Println("====================================================")
}

func RebuildDSN(username, password, conn string) string {
	URL, err := url.Parse(conn)
	if err != nil {
		fmt.Println("error : ", err)
	}

	rebuild := mysql.NewConfig()
	rebuild.User = username
	rebuild.Passwd = password
	rebuild.Net = URL.Scheme
	rebuild.Addr = URL.Host

	return rebuild.FormatDSN()
}
