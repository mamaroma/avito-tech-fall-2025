package postgres

import (
	"avito-tech-fall-2025/internal/domain"
	"avito-tech-fall-2025/internal/repository"

	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TeamRepo struct {
	db *pgxpool.Pool
}

func NewTeamRepository(db *pgxpool.Pool) repository.TeamRepository {
	return &TeamRepo{db: db}
}

func (r *TeamRepo) Create(name string) (int64, error) {
	var id int64
	err := r.db.QueryRow(context.Background(),
		`INSERT INTO teams (name) VALUES ($1) RETURNING id`,
		name).Scan(&id)
	return id, err
}

func (r *TeamRepo) GetByID(id int64) (*domain.Team, error) {
	row := r.db.QueryRow(context.Background(),
		`SELECT id, name FROM teams WHERE id=$1`, id)

	var t domain.Team
	err := row.Scan(&t.ID, &t.Name)
	if errors.Is(err, context.Canceled) {
		return nil, domain.ErrTeamNotFound
	}
	return &t, err
}
