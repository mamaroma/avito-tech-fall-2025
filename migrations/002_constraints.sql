BEGIN;

CREATE INDEX IF NOT EXISTS idx_users_team_active
    ON users (team_id, is_active);

CREATE INDEX IF NOT EXISTS idx_pr_reviewers_user
    ON pr_reviewers (user_id);

CREATE INDEX IF NOT EXISTS idx_pull_requests_author
    ON pull_requests (author_id);

COMMIT;