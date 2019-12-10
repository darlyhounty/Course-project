package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:darlyhounty@tcp:(127.0.0.1:3306)/mytest?charset=utf8")
	fmt.Println(db)
	fmt.Println(err)
	if err != nil {
		fmt.Println("Connection failed")
		return
	}
	fmt.Println("connection succeeded")
	db.Close()

}