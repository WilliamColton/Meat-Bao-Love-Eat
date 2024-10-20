package main

import (
	"Score_System/controller"
	"Score_System/dao"
	"Score_System/db"
	"Score_System/middleware"
	"Score_System/router"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	databaseName := viper.GetString("database_path")
	jwtSecretKey := viper.GetString("jwt_secret_key")
	fmt.Println("databaseName:")
	fmt.Println(databaseName)
	fmt.Println("jwtSecretKey:")
	fmt.Println(jwtSecretKey)

	jwt := middleware.NewJWT(jwtSecretKey)

	DB, err := db.InitDB(databaseName)
	if err != nil {
		return
	}
	d := dao.InitDao(DB)
	c := controller.NewController(jwt, d)
	r := router.NewRouter(c)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
