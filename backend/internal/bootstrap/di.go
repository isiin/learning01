package bootstrap

import (
	"react-ts/backend/internal/domain"
	"react-ts/backend/internal/repo"
	"react-ts/backend/internal/usecase"
)

type Components struct {
	SampleUC   domain.SamplesUseCase
	SampleRepo domain.SampleRepository
}

func NewComponents() *Components {
	sampleRepo := repo.NewSamplesRepository()
	sampleUC := usecase.NewSamplesUseCase(sampleRepo)
	return &Components{
		SampleRepo: sampleRepo,
		SampleUC:   sampleUC,
	}
}
