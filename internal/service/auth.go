package service

import (
	"backend/config"
	"backend/internal/repository"
	"backend/pkg/lib"
	"backend/pkg/util"
	"fmt"
	"time"
)

type IAuthService interface {
	GoogleOAuth(code, pathUrl string) (*string, error)
}

type authService struct {
	userRepository repository.IUserRepository
}

func NewAuthService(userRepository repository.IUserRepository) IAuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (s *authService) GoogleOAuth(code, pathUrl string) (*string, error) {
	tokenRes, err := lib.GetGoogleOauthToken(code)

	if err != nil {
		return nil, err
	}

	user, err := lib.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		return nil, err
	}

	userDB, err := s.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	// Generate JWT Token
	ttl, _ := time.ParseDuration(config.C.AccessTokenExpiresIn)
	fmt.Println(config.C.AccessTokenPrivateKey)
	accessToken, err := util.CreateToken(ttl, userDB.Email, config.C.AccessTokenPrivateKey)

	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}
