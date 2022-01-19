package models

import (
	"context"
	"database/sql"
	"fmt"
)

var (
	db  *sql.DB
	ctx context.Context
)

type GD struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Address string `json:"Address"`
}

func GetAllStores() []GD {
	var arr []GD
	Stores, err := db.QueryContext(ctx, "SELECT * FROM GD.Stores")
	if err != nil {
		fmt.Println("Can't execute query because of" + err.Error())
		return arr
	}
	
	defer Stores.Close()
	for Stores.Next() {
		var name, add string
		var id int
		err := Stores.Scan(&id, &name, &add)
		if err != nil {
			fmt.Println("ERROR")
			continue
		}
		gd := GD{
			Id:      id,
			Name:    name,
			Address: add,
		}
		arr = append(arr, gd)

	}
	return arr
}
