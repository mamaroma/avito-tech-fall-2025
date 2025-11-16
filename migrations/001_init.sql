BEGIN;

CREATE TABLE IF NOT EXISTS teams (
                                     id      BIGSERIAL PRIMARY KEY,
                                     name    TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS users (
                                     id        BIGSERIAL PRIMARY KEY,
                                     name      TEXT NOT NULL,
                                     is_active BOOLEAN NOT NULL DEFAULT TRUE,
                                     team_id   BIGINT NOT NULL REFERENCES teams(id)
    );

CREATE TABLE IF NOT EXISTS pull_requests (
                                             id        BIGSERIAL PRIMARY KEY,
                                             title     TEXT NOT NULL,
                                             author_id BIGINT NOT NULL REFERENCES users(id),
    status    TEXT NOT NULL DEFAULT 'OPEN'
    );

CREATE TABLE IF NOT EXISTS pr_reviewers (
                                            pr_id   BIGINT NOT NULL REFERENCES pull_requests(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id),
    PRIMARY KEY (pr_id, user_id)
    );

COMMIT;