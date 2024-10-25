package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/ssr0016/ecommmerse-app/internal/domain"
	"github.com/ssr0016/ecommmerse-app/internal/dto"
	"github.com/ssr0016/ecommmerse-app/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) Signup(input dto.UserSignup) (string, error) {

	log.Println(input)

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})

	// generate token
	log.Println(user)

	userInfo := fmt.Sprintf("%+v, %+v, %v", user.ID, user.Email, user.UserType)

	// call db to create user
	return userInfo, err
}

func (s *UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUser(email)

	return &user, err
}

func (s *UserService) Login(email, password string) (string, error) {

	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user doest not exist with the provided mail id")
	}

	// compare password and generate token

	return user.Email, nil
}

func (s *UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s *UserService) VerifyCode(id uint, code int) error {
	return nil
}

func (s *UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s *UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s *UserService) UpdateProfile(id uint, input any) error {
	return nil
}

func (s *UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s *UserService) FindCard(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s *UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s *UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s *UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s *UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) {
	return nil, nil
}
