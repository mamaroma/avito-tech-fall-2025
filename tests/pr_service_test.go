package tests

import (
	"testing"

	"avito-tech-fall-2025/internal/app/services"
)

func TestCreatePR_AssignsReviewers(t *testing.T) {
	users := NewMockUserRepo()
	teams := NewMockTeamRepo()
	prs := NewMockPRRepo()

	teamID, _ := teams.Create("backend")

	author, _ := users.Create("Alice", true, teamID)
	_, _ = users.Create("Bob", true, teamID)
	_, _ = users.Create("Charlie", true, teamID)

	svc := services.NewPRService(prs, users, teams)

	id, err := svc.CreatePR("Feature X", author)
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}

	pr, _ := prs.GetByID(id)

	if pr.AuthorID == pr.Reviewers[0] {
		t.Fatalf("author must not be reviewer")
	}
}
