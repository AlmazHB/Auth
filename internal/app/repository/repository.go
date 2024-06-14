package repository

import (
	"crypto/rand"

	"github.com/jmoiron/sqlx"
)

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

// func (a *AuthDB) InsertUser(user models.User) error {
// 	_, err := a.db.Collection("users").InsertOne(context.Background(), user)
// 	if err != nil {
// 		logrus.Errorf("Error inserting user: %s", err.Error())
// 	}
// 	return err
// }

// func (a *AuthDB) FindUserByEmail(email string) (models.User, error) {
// 	var result models.User
// 	err := a.db.Collection("users").FindOne(context.Background(), bson.M{"email": email}).Decode(&result)
// 	if err != nil {
// 		logrus.Errorf("Error finding user by email: %s", err.Error())
// 	}
// 	return result, err
// }

// Функция для генерации случайной соли заданной длины
func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// func hashRefreshToken(refreshToken string, salt []byte) (string, error) {
// 	refreshTokenWithSalt := []byte(refreshToken + string(salt))
// 	hashedRefreshTokenByte, err := bcrypt.GenerateFromPassword([]byte(refreshTokenWithSalt), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "",err
// 	}
// 	hashString := base64.URLEncoding.EncodeToString(hashedRefreshTokenByte)
// 	return hashString, nil
// }
