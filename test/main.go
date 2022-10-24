package test

import (
	"GApp/AuthServer/db"
	"fmt"
)

func TestMain() {
	var err error = db.ConnectDb()
	if err != nil {
		fmt.Println("Can't connect to DB")
		return
	} else {
		fmt.Println("Connected to DB")
	}
}
