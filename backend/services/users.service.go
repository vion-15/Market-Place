package services

import (
	"backend/models"
	"backend/repositories"
	"backend/utils"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req RegisterRequest) (*UserResponse, error)
	Login(req LoginRequest) (*UserResponse, error)
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
	Token string `json:"token"`
	Role  string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
		Role:  newUser.Role,
	}, nil
}

func (s *userService) Login(user LoginRequest) (*UserResponse, error) {
	Email := user.Email
	Password := user.Password

	Email = strings.ToLower(Email)

	userData, err := s.repo.FindByEmail(Email)

	if err != nil {
		fmt.Println("GAGAL DI DB:", err)
		return nil, errors.New("Email atau Password Salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(Password))

	if err != nil {
		fmt.Println("GAGAL DI BCRYPT:", err)
		return nil, errors.New("Email atau Password Salah")
	}

	tokenGenerate, err := utils.GenerateToken(userData.ID, userData.Role)

	if err != nil {
		return nil, errors.New("Gagal Membuat Token")
	}

	return &UserResponse{
		Email: user.Email,
		Token: tokenGenerate,
	}, nil
}

func NewUserService(repo repositories.UserRepositories) *userService {
	return &userService{
		repo: repo,
	}
}
