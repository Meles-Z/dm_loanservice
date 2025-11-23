-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE account_lock_rules (
    id VARCHAR(36) PRIMARY KEY,
    account_id VARCHAR(36) NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    account_status VARCHAR(50) NOT NULL,        -- e.g. 'SECURITISED'
    field_name VARCHAR(100) NOT NULL,        -- e.g. 'interest_rate'
    lock_type VARCHAR(30) NOT NULL DEFAULT 'soft', -- e.g. 'hard', 'soft' 
    lock_reason TEXT NOT NULL,               -- human explanation
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
