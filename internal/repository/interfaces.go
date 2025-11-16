package repository

import "avito-tech-fall-2025/internal/domain"

type UserRepository interface {
	Create(name string, isActive bool, teamID int64) (int64, error)
	GetByID(id int64) (*domain.User, error)
	GetActiveUsersByTeam(teamID int64, excludeID int64) ([]domain.User, error)
	SetActive(id int64, active bool) error
}

type TeamRepository interface {
	Create(name string) (int64, error)
	GetByID(id int64) (*domain.Team, error)
}

type PRRepository interface {
	Create(title string, authorID int64, reviewers []int64) (int64, error)
	GetByID(id int64) (*domain.PullRequest, error)
	UpdateReviewers(id int64, reviewers []int64) error
	SetMerged(id int64) error
	GetByReviewer(userID int64) ([]domain.PullRequest, error)
}
