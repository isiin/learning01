package repository

import (
	"react-ts/backend/internal/domain"
	"react-ts/backend/internal/errs"
)

// TODO repositoryの実装

func NewSurveyRepository() domain.SurveyRepository {
	return &surveyRepository{}
}

type surveyRepository struct {
}

func (r *surveyRepository) GetSurveyors(filter domain.SurveyorFilter) (domain.Surveyors, error) {
	return nil, errs.NewBusinessError(errs.Exclusion, "xxxx")
}
