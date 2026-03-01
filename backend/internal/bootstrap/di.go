package bootstrap

import (
	"react-ts/backend/internal/domain"
	"react-ts/backend/internal/repository"
	"react-ts/backend/internal/usecase"
)

type Components struct {
	SampleUC   domain.SamplesUseCase
	SampleRepo domain.SampleRepository
}

func NewComponents() *Components {
	sampleRepo := repository.NewSamplesRepository()
	sampleUC := usecase.NewSamplesUseCase(sampleRepo)
	return &Components{
		SampleRepo: sampleRepo,
		SampleUC:   sampleUC,
	}
}
