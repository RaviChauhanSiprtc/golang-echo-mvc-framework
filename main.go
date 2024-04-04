package main

import (
	"golang-echo-mvc-framework/routes"
	"golang-echo-mvc-framework/settings"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("settings/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Database
	DB := settings.DatabaseConfig.GetDatabaseConfig(settings.DatabaseConfig{})
	defer DB.Close()

	// Starting server
	echo := routes.Routing.GetRoutes(routes.Routing{})
	err := echo.Start(":" + viper.GetString("port.router"))
	if err != nil {
		log.Fatal(err)
	}

}
