package certificates

type Certificate struct {
	ID         int    `json:"id, omitempty"`
	ExternalID string `json:"external_id"`
	Number     string `json:"number"`
	FullName   string `json:"full_name"`
	Event      string `json:"event"`
	DateAt     string `json:"date_at"`
}
