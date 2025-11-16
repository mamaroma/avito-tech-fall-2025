package postgres

import (
	"avito-tech-fall-2025/internal/domain"
	"avito-tech-fall-2025/internal/repository"

	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PRRepo struct {
	db *pgxpool.Pool
}

func NewPRRepository(db *pgxpool.Pool) repository.PRRepository {
	return &PRRepo{db: db}
}

func (r *PRRepo) Create(title string, authorID int64, reviewers []int64) (int64, error) {
	var id int64
	err := r.db.QueryRow(context.Background(),
		`INSERT INTO pull_requests (title, author_id, status) 
         VALUES ($1, $2, 'OPEN') RETURNING id`,
		title, authorID).Scan(&id)
	if err != nil {
		return 0, err
	}

	for _, rv := range reviewers {
		_, err = r.db.Exec(context.Background(),
			`INSERT INTO pr_reviewers (pr_id, user_id) VALUES ($1, $2)`,
			id, rv)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *PRRepo) GetByID(id int64) (*domain.PullRequest, error) {
	row := r.db.QueryRow(context.Background(),
		`SELECT id, title, author_id, status FROM pull_requests WHERE id=$1`, id)

	var pr domain.PullRequest
	err := row.Scan(&pr.ID, &pr.Title, &pr.AuthorID, &pr.Status)
	if err != nil {
		return nil, domain.ErrPRNotFound
	}

	rows, err := r.db.Query(context.Background(),
		`SELECT user_id FROM pr_reviewers WHERE pr_id=$1`, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var uid int64
		rows.Scan(&uid)
		pr.Reviewers = append(pr.Reviewers, uid)
	}

	return &pr, nil
}

func (r *PRRepo) UpdateReviewers(id int64, reviewers []int64) error {
	_, err := r.db.Exec(context.Background(),
		`DELETE FROM pr_reviewers WHERE pr_id=$1`, id)
	if err != nil {
		return err
	}

	for _, rID := range reviewers {
		_, err = r.db.Exec(context.Background(),
			`INSERT INTO pr_reviewers (pr_id, user_id) VALUES ($1, $2)`,
			id, rID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *PRRepo) SetMerged(id int64) error {
	_, err := r.db.Exec(context.Background(),
		`UPDATE pull_requests SET status='MERGED' WHERE id=$1`, id)
	return err
}

func (r *PRRepo) GetByReviewer(userID int64) ([]domain.PullRequest, error) {
	rows, err := r.db.Query(context.Background(),
		`SELECT pr.id, pr.title, pr.author_id, pr.status
		 FROM pull_requests pr
		 JOIN pr_reviewers rv ON rv.pr_id = pr.id
		 WHERE rv.user_id=$1`, userID)
	if err != nil {
		return nil, err
	}

	var list []domain.PullRequest
	for rows.Next() {
		var pr domain.PullRequest
		rows.Scan(&pr.ID, &pr.Title, &pr.AuthorID, &pr.Status)
		list = append(list, pr)
	}

	return list, nil
}
