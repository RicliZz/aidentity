package authService

import "github.com/RicliZz/aidentity/internal/repositories"

type AuthenticationService struct {
	AuthenticationRepository repositories.AuthenticationRepositoryInterface
}

func NewAuthenticationService(AuthenticationRepository repositories.AuthenticationRepositoryInterface) *AuthenticationService {
	return &AuthenticationService{
		AuthenticationRepository: AuthenticationRepository,
	}
}
