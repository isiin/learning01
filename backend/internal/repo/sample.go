package repo

import (
	"react-ts/backend/internal/domain"
	"react-ts/backend/internal/errs"
)

func NewSamplesRepository() domain.SampleRepository {
	return &repository{}
}

type repository struct {
}

func (r *repository) GetSamples() (domain.Samples, error) {
	return nil, errs.NewBusinessError(errs.Exclusion, "xxxx")
}
