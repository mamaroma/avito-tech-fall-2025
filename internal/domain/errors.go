package domain

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrTeamNotFound = errors.New("team not found")
	ErrPRNotFound   = errors.New("pull request not found")
	ErrPRMerged     = errors.New("pull request already merged")
	ErrNoCandidates = errors.New("no available candidates")
	ErrReviewerSame = errors.New("reviewer cannot be replaced with themselves")
)
