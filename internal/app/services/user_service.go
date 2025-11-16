package services

import (
	"avito-tech-fall-2025/internal/domain"
	"avito-tech-fall-2025/internal/repository"
)

type UserService struct {
	users repository.UserRepository
	teams repository.TeamRepository
}

func NewUserService(users repository.UserRepository, teams repository.TeamRepository) *UserService {
	return &UserService{
		users: users,
		teams: teams,
	}
}

func (s *UserService) CreateUser(name string, isActive bool, teamID int64) (int64, error) {
	_, err := s.teams.GetByID(teamID)
	if err != nil {
		return 0, domain.ErrTeamNotFound
	}
	return s.users.Create(name, isActive, teamID)
}

func (s *UserService) SetActive(id int64, active bool) error {
	_, err := s.users.GetByID(id)
	if err != nil {
		return err
	}
	return s.users.SetActive(id, active)
}

func (s *UserService) Get(id int64) (*domain.User, error) {
	return s.users.GetByID(id)
}
