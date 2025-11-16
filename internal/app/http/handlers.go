package http

import (
	"avito-tech-fall-2025/internal/app/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Handlers struct {
	users *services.UserService
	teams *services.TeamService
	prs   *services.PRService
}

func NewHandlers(u *services.UserService, t *services.TeamService, p *services.PRService) *Handlers {
	return &Handlers{
		users: u,
		teams: t,
		prs:   p,
	}
}

func writeJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func parseID(p httprouter.Params, name string) (int64, error) {
	return strconv.ParseInt(p.ByName(name), 10, 64)
}

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req struct {
		Name     string `json:"name"`
		IsActive bool   `json:"isActive"`
		TeamID   int64  `json:"teamId"`
	}

	_ = json.NewDecoder(r.Body).Decode(&req)

	id, err := h.users.CreateUser(req.Name, req.IsActive, req.TeamID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]int64{"id": id})
}

func (h *Handlers) SetActive(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := parseID(ps, "id")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, nil)
		return
	}

	var req struct {
		Active bool `json:"active"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)

	if err := h.users.SetActive(id, req.Active); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (h *Handlers) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := parseID(ps, "id")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, nil)
		return
	}

	u, err := h.users.Get(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}

	writeJSON(w, http.StatusOK, u)
}

func (h *Handlers) CreateTeam(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req struct {
		Name string `json:"name"`
	}

	_ = json.NewDecoder(r.Body).Decode(&req)

	id, err := h.teams.CreateTeam(req.Name)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]int64{"id": id})
}

func (h *Handlers) GetTeam(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := parseID(ps, "id")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, nil)
		return
	}

	t, err := h.teams.Get(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func (h *Handlers) CreatePR(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req struct {
		Title    string `json:"title"`
		AuthorID int64  `json:"authorId"`
	}

	_ = json.NewDecoder(r.Body).Decode(&req)

	id, err := h.prs.CreatePR(req.Title, req.AuthorID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]int64{"id": id})
}

func (h *Handlers) ReplaceReviewer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := parseID(ps, "id")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, nil)
		return
	}

	oldID, err := parseID(ps, "old")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, nil)
		return
	}

	err = h.prs.ReplaceReviewer(id, oldID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (h *Handlers) MergePR(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := parseID(ps, "id")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, nil)
		return
	}

	pr, err := h.prs.Merge(id)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, pr)
}

func (h *Handlers) GetByReviewer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := parseID(ps, "id")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, nil)
		return
	}

	list, err := h.prs.GetByReviewer(id)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, list)
}
