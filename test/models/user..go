package test

import (
	"GApp/AuthServer/db"
	"GApp/AuthServer/models"
	"context"
)

func ResetAndPopulateSampleUsers() error {
	var err error = db.DB.ResetModel(context.TODO(), (*models.User)(nil))
	return err
}
