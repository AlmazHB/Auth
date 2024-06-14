package app

import (
	"fmt"
	"log"

	"github.com/AlmazHb/Auth/internal/app/configs"
	"github.com/AlmazHb/Auth/internal/app/dbase"
	"github.com/AlmazHb/Auth/internal/app/handler"
	"github.com/AlmazHb/Auth/internal/app/repository"
	"github.com/AlmazHb/Auth/internal/app/service"
	"github.com/AlmazHb/Auth/internal/app/web"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	//d *dbase.DataBase
	r *repository.Repository
	s *service.Service
	h *handler.Handler
	w *web.Server
}

const (
	confPath = "configs"
	confName = "configs"
)

func New() (*App, error) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := configs.Init(confPath, confName); err != nil {
		logrus.Fatalf("Error initializing configs: %s", err.Error())
	}

	a := &App{}
	db, err := dbase.NewDB(dbase.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: "almaz.1",
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatal(err)
	}
	a.r = repository.NewRepository(db)
	a.s = service.NewService(a.r)
	a.h = handler.NewHandler(a.s)

	return a, nil

}

func (a *App) Run() error {
	fmt.Println("Server runing")
	err := a.w.Run("8080", a.h.InitRoutes())
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
