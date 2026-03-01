package bootstrap

import (
	"react-ts/backend/internal/domain"
	"react-ts/backend/internal/repository"
	"react-ts/backend/internal/usecase"
)

type Components struct {
	SampleUC   domain.SamplesUseCase
	SampleRepo domain.SampleRepository
	SurveyUC   domain.SurveyUseCase
	SurveyRepo domain.SurveyRepository
}

func NewComponents() *Components {
	sampleRepo := repository.NewSamplesRepository()
	sampleUC := usecase.NewSamplesUseCase(sampleRepo)
	surveyRepo := repository.NewSurveyRepository()
	surveyUC := usecase.NewSurveyUseCase(surveyRepo)
	return &Components{
		SampleRepo: sampleRepo,
		SampleUC:   sampleUC,
		SurveyRepo: surveyRepo,
		SurveyUC:   surveyUC,
	}
}
