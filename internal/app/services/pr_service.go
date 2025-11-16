package services

import (
	"math/rand"

	"avito-tech-fall-2025/internal/domain"
	"avito-tech-fall-2025/internal/repository"
)

type PRService struct {
	prs   repository.PRRepository
	users repository.UserRepository
	teams repository.TeamRepository
}

func NewPRService(prs repository.PRRepository, users repository.UserRepository, teams repository.TeamRepository) *PRService {
	return &PRService{
		prs:   prs,
		users: users,
		teams: teams,
	}
}

func pickRandom(users []domain.User, count int) []int64 {
	if len(users) == 0 {
		return nil
	}
	if count > len(users) {
		count = len(users)
	}
	perm := rand.Perm(len(users))
	var result []int64
	for i := 0; i < count; i++ {
		result = append(result, users[perm[i]].ID)
	}
	return result
}

func uniqueReviewers(ids []int64) []int64 {
	if len(ids) == 0 {
		return ids
	}
	seen := make(map[int64]struct{}, len(ids))
	res := make([]int64, 0, len(ids))
	for _, id := range ids {
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		res = append(res, id)
	}
	return res
}

func (s *PRService) CreatePR(title string, authorID int64) (int64, error) {
	author, err := s.users.GetByID(authorID)
	if err != nil {
		return 0, domain.ErrUserNotFound
	}

	activeCandidates, err := s.users.GetActiveUsersByTeam(author.TeamID, authorID)
	if err != nil {
		return 0, err
	}

	reviewers := pickRandom(activeCandidates, 2)

	return s.prs.Create(title, authorID, reviewers)
}

func (s *PRService) ReplaceReviewer(prID int64, oldReviewerID int64) error {
	pr, err := s.prs.GetByID(prID)
	if err != nil {
		return err
	}
	if pr.Status == domain.PRMerged {
		return domain.ErrPRMerged
	}

	found := false
	for _, r := range pr.Reviewers {
		if r == oldReviewerID {
			found = true
			break
		}
	}
	if !found {
		return domain.ErrReviewerSame
	}

	oldReviewer, err := s.users.GetByID(oldReviewerID)
	if err != nil {
		return err
	}

	candidates, err := s.users.GetActiveUsersByTeam(oldReviewer.TeamID, oldReviewerID)
	if err != nil {
		return err
	}

	filtered := make([]domain.User, 0, len(candidates))
	for _, c := range candidates {
		if c.ID == pr.AuthorID {
			continue
		}
		filtered = append(filtered, c)
	}

	if len(filtered) == 0 {
		return domain.ErrNoCandidates
	}

	newReviewer := pickRandom(filtered, 1)[0]

	for i := range pr.Reviewers {
		if pr.Reviewers[i] == oldReviewerID {
			pr.Reviewers[i] = newReviewer
			break
		}
	}

	pr.Reviewers = uniqueReviewers(pr.Reviewers)

	return s.prs.UpdateReviewers(prID, pr.Reviewers)
}

func (s *PRService) Merge(prID int64) (*domain.PullRequest, error) {
	pr, err := s.prs.GetByID(prID)
	if err != nil {
		return nil, err
	}
	if pr.Status == domain.PRMerged {
		return pr, nil
	}

	if err := s.prs.SetMerged(prID); err != nil {
		return nil, err
	}

	return s.prs.GetByID(prID)
}

func (s *PRService) GetByReviewer(id int64) ([]domain.PullRequest, error) {
	return s.prs.GetByReviewer(id)
}
