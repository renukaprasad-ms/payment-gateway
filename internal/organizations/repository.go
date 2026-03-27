package organizations

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateOrganization(ctx context.Context, org *Organization) error {

	query := `
	INSERT INTO organizations (name, email, status)
	VALUES ($1, $2, $3)
	RETURNING id, created_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		org.Name,
		org.Email,
		org.Status,
	).Scan(
		&org.ID,
		&org.CreatedAt,
	)
}

func (r *Repository) GetOrganizationByID(ctx context.Context, id string) (*Organization, error) {
	query := `SELECT id,name,email,status,created_at
	FROM organizations
	WHERE id=$1`

	var org Organization

	err := r.db.QueryRow(ctx, query, id).Scan(
		&org.ID,
		&org.Name,
		&org.Email,
		&org.Status,
		&org.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &org, nil
}

func (r *Repository) ListOrganizations(ctx context.Context) ([]Organization, error) {

	query := `
	SELECT id, name, email, status, created_at
	FROM organizations
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orgs []Organization

	for rows.Next() {

		var org Organization

		err := rows.Scan(
			&org.ID,
			&org.Name,
			&org.Email,
			&org.Status,
			&org.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		orgs = append(orgs, org)
	}

	return orgs, nil
}

func (r *Repository) DeleteOrganization(ctx context.Context, id string) error {

	query := `DELETE FROM organizations WHERE id=$1`

	_, err := r.db.Exec(ctx, query, id)

	return err
}
