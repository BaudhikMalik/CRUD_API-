package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB){
	db, err := sql.Open("mysql", "root:@/GD")
	if err != nil {
		fmt.Println("error in getting databse, " + err.Error())
		//panic(fmt.Println(")
	}
	//fmt.Println("database working")
	return db

}
