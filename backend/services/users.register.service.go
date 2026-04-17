package services

import (
	"backend/models"
	"backend/repositories"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req RegisterRequest) (*UserResponse, error)
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type userService struct {
	repo repositories.UserRepositories
}

func (s *userService) Register(user RegisterRequest) (*UserResponse, error) {

	// mengubah Email ke Lowercase
	Email := user.Email
	Email = strings.ToLower(Email)

	// Mengubah nomer telp agar selalu +62
	NoTelp := user.Phone
	if strings.HasPrefix(NoTelp, "0") {
		NoTelp = "+62" + NoTelp[1:]
	}

	// pengecekan email
	_, errEmail := s.repo.FindByEmail(Email)
	if errEmail == nil {
		return &UserResponse{}, errors.New("Email Sudah Digunakan")
	}

	// pengecekan noTelp
	_, errNoTelp := s.repo.FindByNoTelp(NoTelp)
	if errNoTelp == nil {
		return &UserResponse{}, errors.New("Nomer Telphone Sudah Digunakan")
	}

	//Hashing Password
	Password := user.Password
	result, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return &UserResponse{}, errors.New("Gagal memproses password")
	}
	HashedPassword := string(result)

	// mapping ke models
	newUser := &models.Users{
		Name:            user.Name,
		Email:           Email,
		Phone:           NoTelp,
		Password:        HashedPassword,
		IsEmailVerified: false,
		IsPhoneVerified: false,
		Role:            "buyer",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	errCreate := s.repo.Create(newUser)
	if errCreate != nil {
		return &UserResponse{}, errors.New("Terjadi Kesalahan System")
	}

	return &UserResponse{
		Name:  user.Name,
		Email: Email,
		Phone: NoTelp,
	}, nil
}

func NewUserService(repo repositories.UserRepositories) *userService {
	return &userService{
		repo: repo,
	}
}
