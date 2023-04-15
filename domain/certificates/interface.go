package certificates

import "context"

type CertificateRepository interface {
	List(context.Context) ([]Certificate, error)
	GetByExternalIDAndEvent(ctx context.Context, externalID string) (*Certificate, error)
	StoreCert(ctx context.Context, cert Certificate) error
}

type CertificateService interface {
	// List(context.Context) ([]Certificate, error)
	GetByExternalIDAndEvent(ctx context.Context, externalID string) (*CertificateDTO, error)
}
