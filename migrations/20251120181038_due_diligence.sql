-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE DueDiligenceChecklistItems (
    id VARCHAR(30) PRIMARY KEY,
    account_id VARCHAR(30) NOT NULL REFERENCES accounts(id),
    checklist_item VARCHAR(255) NOT NULL,
    status VARCHAR(20) NOT NULL, -- 'pending', 'pass', 'fail'
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
