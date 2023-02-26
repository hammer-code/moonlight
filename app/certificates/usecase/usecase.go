package usecase

import (
	"context"

	"github.com/hammer-code/moonlight/domain/certificates"
)

type usecase struct {
	repo certificates.CertificateRepository
}

func Newusecase(repo certificates.CertificateRepository) certificates.CertificateService {
	return &usecase{
		repo: repo,
	}
}

func (u usecase) GetByExternalIDAndEvent(ctx context.Context, externalID string, event string) (*certificates.CertificateDTO, error) {
	cer, err := u.repo.GetByExternalIDAndEvent(ctx, externalID, event)
	if err != nil {
		return nil, err
	}

	return certificates.ToDTO(*cer), nil
}
