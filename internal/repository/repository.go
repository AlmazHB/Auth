package repository

import "github.com/jmoiron/sqlx"

type AuthDB struct {
	db *sqlx.DB
}

func NewAuthDB(db *sqlx.DB) *AuthDB {
	return &AuthDB{db: db}
}

type Repository struct {
	AuthDB *AuthDB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthDB: NewAuthDB(db),
	}
}
