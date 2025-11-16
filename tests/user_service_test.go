package tests

import (
	"testing"

	"avito-tech-fall-2025/internal/app/services"
)

func TestUserService_CreateUser(t *testing.T) {
	users := NewMockUserRepo()
	teams := NewMockTeamRepo()
	svc := services.NewUserService(users, teams)

	teamID, _ := teams.Create("backend")

	id, err := svc.CreateUser("Alice", true, teamID)
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}

	u, _ := users.GetByID(id)
	if u.Name != "Alice" {
		t.Fatalf("bad name")
	}
}
