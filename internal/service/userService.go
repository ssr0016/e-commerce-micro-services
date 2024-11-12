package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/ssr0016/ecommmerse-app/config"
	"github.com/ssr0016/ecommmerse-app/internal/domain"
	"github.com/ssr0016/ecommmerse-app/internal/dto"
	"github.com/ssr0016/ecommmerse-app/internal/helper"
	"github.com/ssr0016/ecommmerse-app/internal/repository"
	"github.com/ssr0016/ecommmerse-app/pkg/notification"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s *UserService) Signup(input dto.UserSignup) (string, error) {
	hashedPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	})

	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s *UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUser(email)

	return &user, err
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}

	err = s.Auth.VerifyPassword(user.Password, password)
	if err != nil {
		return "", err
	}

	// generate token
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s *UserService) isVerified(id uint) bool {
	currentUser, err := s.Repo.FindUserById(id)

	return err == nil && currentUser.Verified
}

func (s *UserService) GetVerificationCode(e domain.User) error {
	// if user already vefified
	if s.isVerified(e.ID) {
		return errors.New("user already verified")
	}

	// generate vefication code
	code, err := s.Auth.GenerateCode()
	if err != nil {
		return err
	}

	// update user
	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	_, err = s.Repo.UpdateUser(e.ID, user)
	if err != nil {
		return errors.New("unable to update verification code")
	}

	user, _ = s.Repo.FindUserById(e.ID)

	// send SMS
	notificationClient := notification.NewNotificationClient(s.Config)

	msg := fmt.Sprintf("Your verification code is %v", code)

	err = notificationClient.SendSMS(user.Phone, msg)
	if err != nil {
		return errors.New("error on sending sms")
	}

	//return verification code
	return nil
}

func (s *UserService) VerifyCode(id uint, code int) error {
	// if user already vefified
	if s.isVerified(id) {
		return errors.New("user already verified")
	}

	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("verification code does not match")
	}

	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code is expired")
	}

	// update user
	updateUser := domain.User{
		Verified: true,
	}

	_, err = s.Repo.UpdateUser(id, updateUser)
	if err != nil {
		return errors.New("unable to verify user")
	}

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

func (s *UserService) BecomeSeller(id uint, input dto.SellerInput) (string, error) {
	// find the existing user
	user, _ := s.Repo.FindUserById(id)
	if user.UserType == domain.SELLER {
		return "", errors.New("you have already joined seller program")
	}

	// update user
	seller, err := s.Repo.UpdateUser(id, domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.PhoneNumber,
		UserType:  domain.SELLER,
	})

	if err != nil {
		return "", err
	}

	// generating token

	token, err := s.Auth.GenerateToken(user.ID, user.Email, seller.UserType)
	if err != nil {
		return "", err
	}

	//  create bank account information

	err = s.Repo.CreateBankAccount(domain.BankAccount{
		BankAccount: input.BankAccountNumber,
		SwiftCode:   input.SwiftCode,
		Payment:     input.PaymentType,
		UserID:      id,
	})

	return token, err
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
