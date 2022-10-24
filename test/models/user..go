package test

import (
	"GApp/AuthServer/db"
	"GApp/AuthServer/models"
	"context"
	"fmt"
)

func ResetAndPopulateSampleUsers() error {
	err := db.DB.ResetModel(context.TODO(), (*models.User)(nil))
	if err != nil {
		fmt.Println("Error creating User table")
		fmt.Println(err)
		return err
	}

	users := []models.User{
		{Name: "User1", Email: "user1@email.com"},
		{Name: "User2", Email: "user2@email.com"},
	}

	_, err = db.DB.NewInsert().Model(&users).Exec(context.TODO())
	if err != nil {
		fmt.Println("Error inserting records to table")
		fmt.Println(err)
		return err
	}
	fmt.Println("User Reset And Populate Successful")
	return nil
}
