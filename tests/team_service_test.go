package tests

import (
	"testing"

	"avito-tech-fall-2025/internal/app/services"
)

func TestTeamService(t *testing.T) {
	teams := NewMockTeamRepo()
	users := NewMockUserRepo()

	svc := services.NewTeamService(teams, users)

	id, _ := svc.CreateTeam("backend")

	tm, _ := svc.Get(id)
	if tm.Name != "backend" {
		t.Fatalf("wrong name")
	}
}
