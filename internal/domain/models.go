package domain

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
	TeamID   int64  `json:"teamId"`
}

type Team struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type PRStatus string

const (
	PROpen   PRStatus = "OPEN"
	PRMerged PRStatus = "MERGED"
)

type PullRequest struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	AuthorID  int64    `json:"authorId"`
	Status    PRStatus `json:"status"`
	Reviewers []int64  `json:"reviewers"`
}
