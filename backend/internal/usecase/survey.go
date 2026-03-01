package usecase

import (
	"react-ts/backend/internal/domain"
)

func NewSurveyUseCase(repo domain.SurveyRepository) domain.SurveyUseCase {
	return &surveyUseCase{
		repo: repo,
	}
}

type surveyUseCase struct {
	repo domain.SurveyRepository
}

func (u *surveyUseCase) GetSurveyors(filter domain.SurveyorFilter) (domain.Surveyors, error) {
	md, err := u.repo.GetSurveyors(filter)
	if err != nil {
		// TODO Errのラップ
		return nil, err
	}
	return md, nil
}
