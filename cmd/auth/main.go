package main

import (
	"log"

	"github.com/AlmazHb/Auth/internal/database"
)

func main() {
	_, err := database.NewDB(database.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "almaz.1",
		DBname:   "auth-db",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}
	// repos :=repository.NewRepository(db)
	// services := ser
	// handlers := handler.NewHandler(services)
	// server := new(web.Server)
	// if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
	// 	logrus.Fatal(err)
	// }

}
