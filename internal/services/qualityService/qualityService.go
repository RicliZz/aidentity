package qualityService

import "github.com/RicliZz/aidentity/internal/repositories"

type QualityService struct {
	QualityRepository repositories.QualityRepositoryInterface
}

func NewQualityService(QualityRepository repositories.QualityRepositoryInterface) *QualityService {
	return &QualityService{
		QualityRepository: QualityRepository,
	}
}
