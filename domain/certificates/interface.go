package certificates

import "context"

type CertificateRepository interface {
	List(context.Context) ([]Certificate, error)
	GetByExternalID(ctx context.Context, externalID string) (*Certificate, error)
}

type CertificateService interface {
	List(context.Context) ([]Certificate, error)
	GetByExternalID(ctx context.Context, externalID string) (*Certificate, error)
}
