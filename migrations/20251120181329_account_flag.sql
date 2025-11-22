-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE account_flags (
    id VARCHAR(36) PRIMARY KEY,
    account_id VARCHAR(36) NOT NULL REFERENCES accounts(id),
    flag_type VARCHAR(100) NOT NULL, -- e.g., 'ineligible', 'high-risk', 'fraud'
    reason TEXT ,
    flagged_by INT NOT NULL,
    flagged_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
