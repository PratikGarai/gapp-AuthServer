package models

import (
	"GApp/AuthServer/db"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID     int64 `bun:",pk,autoincrement"`
	Name   string
	Emails string
}

func ResetAndPopulateSampleUsers() error {
	var err error = db.DB.ResetModel(nil, (*User)(nil))

	return err
}
