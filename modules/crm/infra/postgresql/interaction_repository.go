package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/manab-pr/evtaarpro/modules/crm/domain/entities"
)

type CustomerInteractionRepository struct {
	db *sql.DB
}

func NewCustomerInteractionRepository(db *sql.DB) *CustomerInteractionRepository {
	return &CustomerInteractionRepository{db: db}
}

func (r *CustomerInteractionRepository) Create(ctx context.Context, interaction *entities.CustomerInteraction) error {
	query := `
		INSERT INTO customer_interactions (id, customer_id, user_id, interaction_type,
			subject, notes, interaction_date, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.ExecContext(ctx, query,
		interaction.ID, interaction.CustomerID, interaction.UserID,
		interaction.InteractionType, interaction.Subject, interaction.Notes,
		interaction.InteractionDate, interaction.CreatedAt,
	)
	return err
}

func (r *CustomerInteractionRepository) GetByID(ctx context.Context, id string) (*entities.CustomerInteraction, error) {
	query := `
		SELECT id, customer_id, user_id, interaction_type, subject, notes,
			interaction_date, created_at
		FROM customer_interactions WHERE id = $1
	`
	interaction := &entities.CustomerInteraction{}
	var subject, notes sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&interaction.ID, &interaction.CustomerID, &interaction.UserID,
		&interaction.InteractionType, &subject, &notes,
		&interaction.InteractionDate, &interaction.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("interaction not found")
	}
	if err != nil {
		return nil, err
	}

	if subject.Valid {
		interaction.Subject = &subject.String
	}
	if notes.Valid {
		interaction.Notes = &notes.String
	}

	return interaction, nil
}

func (r *CustomerInteractionRepository) ListByCustomer(ctx context.Context, customerID string, offset, limit int) ([]*entities.CustomerInteraction, int, error) {
	countQuery := `SELECT COUNT(*) FROM customer_interactions WHERE customer_id = $1`
	query := `
		SELECT id, customer_id, user_id, interaction_type, subject, notes,
			interaction_date, created_at
		FROM customer_interactions
		WHERE customer_id = $1
		ORDER BY interaction_date DESC
		LIMIT $2 OFFSET $3
	`

	var total int
	err := r.db.QueryRowContext(ctx, countQuery, customerID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.QueryContext(ctx, query, customerID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	interactions := []*entities.CustomerInteraction{}
	for rows.Next() {
		interaction := &entities.CustomerInteraction{}
		var subject, notes sql.NullString

		err := rows.Scan(
			&interaction.ID, &interaction.CustomerID, &interaction.UserID,
			&interaction.InteractionType, &subject, &notes,
			&interaction.InteractionDate, &interaction.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		if subject.Valid {
			interaction.Subject = &subject.String
		}
		if notes.Valid {
			interaction.Notes = &notes.String
		}

		interactions = append(interactions, interaction)
	}

	return interactions, total, nil
}

func (r *CustomerInteractionRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM customer_interactions WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
