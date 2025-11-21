-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE DueDiligenceChecklistItems (
    id VARCHAR(36) PRIMARY KEY,
    account_id VARCHAR(36) NOT NULL REFERENCES accounts(id),
    checklist_item TEXT NOT NULL,
    status VARCHAR(255) NOT NULL, -- 'pending', 'pass', 'fail'
    comments TEXT,
    created_by INT, -- user id
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
