package repositories

type QualityRepositoryInterface interface {
	CreateQuality(nameNewQuality string) error
}
