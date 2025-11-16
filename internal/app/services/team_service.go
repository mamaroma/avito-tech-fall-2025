package services

import (
	"avito-tech-fall-2025/internal/domain"
	"avito-tech-fall-2025/internal/repository"
)

type TeamService struct {
	teams repository.TeamRepository
	users repository.UserRepository
}

func NewTeamService(teams repository.TeamRepository, users repository.UserRepository) *TeamService {
	return &TeamService{
		teams: teams,
		users: users,
	}
}

func (s *TeamService) CreateTeam(name string) (int64, error) {
	return s.teams.Create(name)
}

func (s *TeamService) Get(id int64) (*domain.Team, error) {
	return s.teams.GetByID(id)
}
