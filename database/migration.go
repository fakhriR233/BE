package database

import (
	"_waysbook/models"
	"_waysbook/pkg/mysql"
	"fmt"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Transaction{},
		&models.Book{},
		&models.Profile{},
		&models.Cart{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
