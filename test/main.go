package test

import (
	"GApp/AuthServer/db"
	test "GApp/AuthServer/test/models"
	"fmt"
)

func TestMain() {
	var err error = db.ConnectDb()
	if err != nil {
		fmt.Println("Can't connect to test DB")
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to test DB")
	test.ResetAndPopulateSampleUsers()
}
