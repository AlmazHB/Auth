package dbase

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLMode  string
}

type DataBase struct {
}

func NewDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// func (dbase *DataBase) InitDB() (*sqlx.DB, error) {
// 	db, err := dbase.NewDB(Config{
// 		Host:     "localhost",
// 		Port:     "5436",
// 		Username: "postgres",
// 		Password: "almaz.1",
// 		DBname:   "auth-db",
// 		SSLMode:  "disable",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db, nil
// }
