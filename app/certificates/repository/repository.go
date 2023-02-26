package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hammer-code/moonlight/domain/certificates"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) certificates.CertificateRepository {
	return &repository{
		db: db,
	}
}

// struct table
// ID         int    `json:"id, omitempty"`
// ExternalID string `json:"external_id"`
// Number     string `json:"number"`
// FullName   string `json:"full_name"`
// Event      string `json:"event"`
// DateAt     string `json:"date_at"`

const (
	tableName    = "certificates"
	publicColumn = "external_id,name,image_link,share_link,event,created_at"
	insertColumn = "external_id,name,image_link,share_link,event"
	// column       = "id,external_id,number,full_name,event,date_at"
)

func (repo repository) List(ctx context.Context) ([]certificates.Certificate, error) {
	selct := fmt.Sprintf("select %s from %s", publicColumn, tableName)
	rows, err := repo.db.QueryContext(ctx, selct)
	if err != nil {
		return nil, err
	}

	var certs []certificates.Certificate
	for rows.Next() {

		var cert certificates.Certificate

		err = rows.Scan(
			&cert.ExternalID,
			&cert.Name,
			&cert.ImageLink,
			&cert.ShareLink,
			&cert.Event,
			&cert.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		certs = append(certs, cert)
	}
	return certs, nil
}

func (repo repository) GetByExternalID(ctx context.Context, externalID string) (*certificates.Certificate, error) {
	selct := fmt.Sprintf("select %s from %s where external_id = %s", publicColumn, tableName, externalID)
	rows, err := repo.db.QueryContext(ctx, selct)
	if err != nil {
		return nil, err
	}

	var cert certificates.Certificate

	if rows.Next() {

		err = rows.Scan(
			&cert.ExternalID,
			&cert.Name,
			&cert.ImageLink,
			&cert.ShareLink,
			&cert.Event,
			&cert.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
	}
	return &cert, nil
}

func (repo repository) StoreCert(ctx context.Context, cert certificates.Certificate) error {
	// insertColumn = "external_id,name,image_link,share_link,event"
	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES($1, $2, $3, $4, $5)", tableName, insertColumn)
	_, err := repo.db.Exec(sql, cert.ExternalID, cert.Name, cert.ImageLink, cert.ShareLink, cert.Event)
	if err != nil {
		return err
	}

	return nil
}
