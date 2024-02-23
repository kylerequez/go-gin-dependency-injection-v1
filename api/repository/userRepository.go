package repository

import (
	"github.com/google/uuid"
	"github.com/kylerequez/go-gin-dependency-injection-v1/types"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetAllUsers() (u *[]types.User, err error) {
	var users []types.User

	result := ur.db.Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (ur *UserRepository) GetUserById(id uuid.UUID) (u *types.User, err error) {
	var user types.User
	result := ur.db.Find(&user, types.User{ID: id})
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return &user, nil
}

func (ur *UserRepository) CreateUser(newUser *types.User) (u *types.User, err error) {
	result := ur.db.Create(newUser)
	if err := result.Error; err != nil || result.RowsAffected <= 0 {
		return nil, err
	}
	return newUser, nil
}

func (ur *UserRepository) UpdateUser(user *types.User) (u *types.User, err error) {
	result := ur.db.Save(user)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return user, nil
}

func (ur *UserRepository) DeleteUser(id uuid.UUID) (res *gorm.DB, err error) {
	result := ur.db.Delete(&types.User{ID: id})
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return result, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (u *types.User, err error) {
	var user types.User
	result := ur.db.Find(&user, types.User{Email: email})
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return &user, nil
}
