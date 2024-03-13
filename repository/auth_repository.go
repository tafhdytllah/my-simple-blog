package repository

import (
	"my-simple-blog/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	EmailExist(email string) bool
	Register(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

// method
func (r *authRepository) EmailExist(email string) bool {
	var user entity.User

	result := r.db.First(&user, "email = ?", email)

	return result.Error == nil
}

// create record
func (r *authRepository) Register(user *entity.User) error {
	result := r.db.Create(&user)

	return result.Error
}

// get user by email
func (r *authRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	result := r.db.First(&user, "email = ?", email)

	return &user, result.Error
}
