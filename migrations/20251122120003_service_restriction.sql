-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE servicing_restrictions (
    id VARCHAR(36) PRIMARY KEY,
    account_id VARCHAR(36) REFERENCES accounts(id) ON DELETE CASCADE,-- NULL loan_id = global rule for all loans
    restriction_type VARCHAR(50) NOT NULL, -- e.g. 'ACTION_BLOCK', 'LIMITED_ACTION', etc
    action_name VARCHAR(100) NOT NULL,-- e.g. 'prepayment', 'waive_late_fee', 'modify_term'
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    reason TEXT NOT NULL,
    source VARCHAR(100) DEFAULT 'system', -- system, compliance, securitisation_pool, etc.
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
