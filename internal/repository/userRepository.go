package repository

import (
	"errors"
	"log"

	"github.com/ssr0016/ecommmerse-app/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)

	CreateBankAccount(entity domain.BankAccount) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		log.Printf("create user error %v\n", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return user, nil
}

func (r *userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		log.Printf("find user error %v\n", err)
		return domain.User{}, errors.New("failed to find user")
	}

	return user, nil
}

func (r *userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, id).Error
	if err != nil {
		log.Printf("find user error %v\n", err)
		return domain.User{}, errors.New("user does not exist")
	}

	return user, nil
}

func (r *userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {

	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).Updates(u).Error

	if err != nil {
		log.Printf("error on update %v\n", err)
		return domain.User{}, errors.New("failed update user")
	}

	return user, nil
}

func (r *userRepository) CreateBankAccount(entity domain.BankAccount) error {

	return r.db.Create(&entity).Error
}
