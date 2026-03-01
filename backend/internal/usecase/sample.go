package usecase

import (
	"react-ts/backend/internal/domain"
)

func NewSamplesUseCase(repo domain.SampleRepository) domain.SamplesUseCase {
	return &sampleUseCase{
		repo: repo,
	}
}

type sampleUseCase struct {
	repo domain.SampleRepository
}

func (u *sampleUseCase) GetSamples() (domain.Samples, error) {
	md, err := u.repo.GetSamples()
	if err != nil {
		// TODO Errのラップ
		return nil, err
	}
	return md, nil
}
