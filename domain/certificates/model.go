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
