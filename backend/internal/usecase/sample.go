package usecase

import (
	"errors"
	"fmt"
	"react-ts/backend/internal/domain"
	"react-ts/backend/internal/errs"
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
	if _, err := u.repo.GetSamples(); err != nil {
		// return err
	}
	cause := errors.New("xxx")
	cause = fmt.Errorf("yyy cause: %w", cause)
	cause = fmt.Errorf("zzz cause: %w", cause)
	return nil, errs.NewSystemError("aaaa", cause)
}
