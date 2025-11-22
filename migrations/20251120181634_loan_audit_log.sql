-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE AccountAuditLog (
    id VARCHAR(36) PRIMARY KEY,
    account_id VARCHAR(36) NOT NULL REFERENCES accounts(id),
    action VARCHAR(100), -- e.g., 'checklist_item_updated', 'loan_flag_added'
    details TEXT,       -- store previous and new values
    performed_by INT,   -- user id
    performed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
