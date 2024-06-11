package test

import (
	"fmt"
	"net/http"

	"mandalart.com/schemas"
	"mandalart.com/utils"
)

func CreateSheet(w http.ResponseWriter, r *http.Request){
	var user schemas.User
	fmt.Println("gogo")
	utils.DB.First(&user, 1)
	utils.DB.Model(&user).Association("Sheets").Append(&schemas.Sheet{})
}