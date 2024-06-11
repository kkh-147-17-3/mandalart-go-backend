package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mandalart.com/schemas"
)

var DB *gorm.DB

func InitDatabase() {
	fmt.Println("start go connect databse")
	var err error
	DB, err = gorm.Open(postgres.Open("user=eggtart password=tkfkdgo486! host=43.203.193.216 port=5432 dbname=eggtart_db_1"))
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(schemas.User{}, schemas.Sheet{}, schemas.Cell{}, schemas.Todo{})
}
