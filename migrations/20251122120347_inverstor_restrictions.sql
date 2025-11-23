-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE investor_restrictions (
    id VARCHAR(36) PRIMARY KEY,
    account_id VARCHAR(36) NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,   
    restriction_scope VARCHAR(20) NOT NULL,-- 'FIELD', 'ACTION'
    field_name VARCHAR(100),-- required if restriction_scope = 'FIELD'
    action_name VARCHAR(100),-- required if restriction_scope = 'ACTION'
    rule_type VARCHAR(30) NOT NULL DEFAULT 'hard',-- 'hard' = never allowed  -- 'approval_required' = request must go to investor
    reason TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
