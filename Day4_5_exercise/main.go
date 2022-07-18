package main

import (
	"Day4_5_exercise/Config"
	"Day4_5_exercise/Models"
	"Day4_5_exercise/Routes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Retailer{})
	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Customer{})
	Config.DB.AutoMigrate(&Models.Order{})
	r := Routes.SetupRouter()
	//running
	err := r.Run()

	if err != nil {
		panic("Error: Facing Error to start the Router")
	}
}
