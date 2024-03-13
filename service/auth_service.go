package service

import (
	"my-simple-blog/dto"
	"my-simple-blog/entity"
	"my-simple-blog/errorhandler"
	"my-simple-blog/helper"
	"my-simple-blog/repository"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

// method
func (s *authService) Register(req *dto.RegisterRequest) error {
	// check email existing
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &errorhandler.BadRequestError{Message: "email already registered"}
	}

	// check match password
	if req.Password != req.PasswordConfirm {
		return &errorhandler.BadRequestError{Message: "password not match"}
	}

	// hash passwrod
	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	// set value object user
	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: passwordHash,
		Gender:   req.Gender,
	}

	// send to register repository
	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	//check user by email
	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "wrong email or password"}
	}

	//check match password
	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.BadRequestError{Message: "wrong email or password"}
	}

	//gemerate token
	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	//set response data
	data = dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}

	return &data, nil
}
