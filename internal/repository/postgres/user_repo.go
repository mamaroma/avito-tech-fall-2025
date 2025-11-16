package postgres

import (
	"avito-tech-fall-2025/internal/domain"
	"avito-tech-fall-2025/internal/repository"

	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) repository.UserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(name string, isActive bool, teamID int64) (int64, error) {
	var id int64
	err := r.db.QueryRow(context.Background(),
		`INSERT INTO users (name, is_active, team_id) VALUES ($1, $2, $3) RETURNING id`,
		name, isActive, teamID).Scan(&id)
	return id, err
}

func (r *UserRepo) GetByID(id int64) (*domain.User, error) {
	row := r.db.QueryRow(context.Background(),
		`SELECT id, name, is_active, team_id FROM users WHERE id=$1`, id)

	var u domain.User
	err := row.Scan(&u.ID, &u.Name, &u.IsActive, &u.TeamID)
	if errors.Is(err, context.Canceled) {
		return nil, domain.ErrUserNotFound
	}
	return &u, err
}

func (r *UserRepo) GetActiveUsersByTeam(teamID int64, excludeID int64) ([]domain.User, error) {
	rows, err := r.db.Query(context.Background(),
		`SELECT id, name, is_active, team_id 
		 FROM users 
		 WHERE team_id=$1 AND is_active=true AND id <> $2`,
		teamID, excludeID)
	if err != nil {
		return nil, err
	}

	var list []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Name, &u.IsActive, &u.TeamID); err != nil {
			return nil, err
		}
		list = append(list, u)
	}

	return list, nil
}

func (r *UserRepo) SetActive(id int64, active bool) error {
	_, err := r.db.Exec(context.Background(),
		`UPDATE users SET is_active=$1 WHERE id=$2`, active, id)
	return err
}
