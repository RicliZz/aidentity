package profileService

import "github.com/RicliZz/aidentity/internal/repositories"

type ProfileService struct {
	ProfileRepository repositories.ProfileRepositoryInterface
}

func NewProfileService(ProfileRepository repositories.ProfileRepositoryInterface) *ProfileService {
	return &ProfileService{
		ProfileRepository: ProfileRepository,
	}
}
