package repository

import (
	"react-ts/backend/internal/domain"
)

// TODO repositoryの実装

func NewSamplesRepository() domain.SampleRepository {
	return &repository{}
}

type repository struct {
}

func (r *repository) GetSamples() (domain.Samples, error) {
	//TODO
	ret := domain.Samples{
		{ID: "01", Name: "サンプル1"},
		{ID: "02", Name: "サンプル2"},
	}
	return ret, nil
}
