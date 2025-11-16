package http

import (
	"net/http"

	"avito-tech-fall-2025/internal/app/services"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	router *httprouter.Router
}

func NewRouter(
	users *services.UserService,
	teams *services.TeamService,
	prs *services.PRService,
) *Server {
	h := NewHandlers(users, teams, prs)

	r := httprouter.New()

	r.POST("/users", h.CreateUser)
	r.PUT("/users/:id/active", h.SetActive)
	r.GET("/users/:id", h.GetUser)

	r.POST("/teams", h.CreateTeam)
	r.GET("/teams/:id", h.GetTeam)

	r.POST("/prs", h.CreatePR)
	r.PUT("/prs/:id/reviewers/:old", h.ReplaceReviewer)
	r.PUT("/prs/:id/merge", h.MergePR)
	r.GET("/prs/reviewer/:id", h.GetByReviewer)

	return &Server{router: r}
}

func (s *Server) Listen(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
