package certificates

import "time"

type Certificate struct {
	ID         int       `json:"id"`
	ExternalID string    `json:"external_id"`
	Name       string    `json:"name"`
	ImageLink  string    `json:"image_link"`
	ShareLink  string    `json:"share_link"`
	Event      string    `json:"event"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CertificateDTO struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	ImageLink string `json:"image_link"`
	ShareLink string `json:"share_link"`
	Event     string `json:"event"`
}

func ToDTO(cer Certificate) *CertificateDTO {
	return &CertificateDTO{
		ID:        cer.ExternalID,
		Name:      cer.Name,
		ImageLink: cer.ImageLink,
		ShareLink: cer.ShareLink,
		Event:     cer.Event,
	}
}
