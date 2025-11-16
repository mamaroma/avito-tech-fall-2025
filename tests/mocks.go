package tests

import "avito-tech-fall-2025/internal/domain"

type MockUserRepo struct {
	Users map[int64]*domain.User
}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{Users: make(map[int64]*domain.User)}
}

func (m *MockUserRepo) Create(name string, isActive bool, teamID int64) (int64, error) {
	id := int64(len(m.Users) + 1)
	m.Users[id] = &domain.User{
		ID:       id,
		Name:     name,
		IsActive: isActive,
		TeamID:   teamID,
	}
	return id, nil
}

func (m *MockUserRepo) GetByID(id int64) (*domain.User, error) {
	u, ok := m.Users[id]
	if !ok {
		return nil, domain.ErrUserNotFound
	}
	return u, nil
}

func (m *MockUserRepo) GetActiveUsersByTeam(teamID int64, excludeID int64) ([]domain.User, error) {
	var res []domain.User
	for _, u := range m.Users {
		if u.TeamID == teamID && u.IsActive && u.ID != excludeID {
			res = append(res, *u)
		}
	}
	return res, nil
}

func (m *MockUserRepo) SetActive(id int64, active bool) error {
	u, ok := m.Users[id]
	if !ok {
		return domain.ErrUserNotFound
	}
	u.IsActive = active
	return nil
}

type MockTeamRepo struct {
	Teams map[int64]*domain.Team
}

func NewMockTeamRepo() *MockTeamRepo {
	return &MockTeamRepo{Teams: make(map[int64]*domain.Team)}
}

func (m *MockTeamRepo) Create(name string) (int64, error) {
	id := int64(len(m.Teams) + 1)
	m.Teams[id] = &domain.Team{
		ID:   id,
		Name: name,
	}
	return id, nil
}

func (m *MockTeamRepo) GetByID(id int64) (*domain.Team, error) {
	t, ok := m.Teams[id]
	if !ok {
		return nil, domain.ErrTeamNotFound
	}
	return t, nil
}

type MockPRRepo struct {
	PRs map[int64]*domain.PullRequest
}

func NewMockPRRepo() *MockPRRepo {
	return &MockPRRepo{PRs: make(map[int64]*domain.PullRequest)}
}

func (m *MockPRRepo) Create(title string, authorID int64, reviewers []int64) (int64, error) {
	id := int64(len(m.PRs) + 1)
	cp := make([]int64, len(reviewers))
	copy(cp, reviewers)
	m.PRs[id] = &domain.PullRequest{
		ID:        id,
		Title:     title,
		AuthorID:  authorID,
		Status:    domain.PROpen,
		Reviewers: cp,
	}
	return id, nil
}

func (m *MockPRRepo) GetByID(id int64) (*domain.PullRequest, error) {
	pr, ok := m.PRs[id]
	if !ok {
		return nil, domain.ErrPRNotFound
	}
	cp := *pr
	if pr.Reviewers != nil {
		cp.Reviewers = make([]int64, len(pr.Reviewers))
		copy(cp.Reviewers, pr.Reviewers)
	}
	return &cp, nil
}

func (m *MockPRRepo) UpdateReviewers(id int64, reviewers []int64) error {
	pr, ok := m.PRs[id]
	if !ok {
		return domain.ErrPRNotFound
	}
	cp := make([]int64, len(reviewers))
	copy(cp, reviewers)
	pr.Reviewers = cp
	return nil
}

func (m *MockPRRepo) SetMerged(id int64) error {
	pr, ok := m.PRs[id]
	if !ok {
		return domain.ErrPRNotFound
	}
	pr.Status = domain.PRMerged
	return nil
}

func (m *MockPRRepo) GetByReviewer(userID int64) ([]domain.PullRequest, error) {
	var res []domain.PullRequest
	for _, pr := range m.PRs {
		for _, r := range pr.Reviewers {
			if r == userID {
				cp := *pr
				if pr.Reviewers != nil {
					cp.Reviewers = make([]int64, len(pr.Reviewers))
					copy(cp.Reviewers, pr.Reviewers)
				}
				res = append(res, cp)
				break
			}
		}
	}
	return res, nil
}
