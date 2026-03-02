package repository

import (
	"react-ts/backend/internal/domain"
)

// TODO repositoryの実装

func NewSurveyRepository() domain.SurveyRepository {
	return &surveyRepository{}
}

type surveyRepository struct {
}

func (r *surveyRepository) GetSurveyors(filter domain.SurveyorFilter) (domain.Surveyors, error) {
	//TODO
	ret := domain.Surveyors{
		{ID: "000001", Name: "調査員1", OfficeID: "XX", OfficeName: "〇〇事業所"},
		{ID: "000002", Name: "調査員2", OfficeID: "XX", OfficeName: "〇〇事業所"},
	}
	return ret, nil
}
