package database

import (
	"Trello/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	dsn := "host=localhost port=5432 user=postgres dbname=trello password=postgres sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.Workspace{},
		&model.User{},
		&model.UserWorkspaceRole{},
		&model.Task{},
		&model.SubTask{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
