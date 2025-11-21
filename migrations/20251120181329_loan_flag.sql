-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE LoanFlags (
    id VARCHAR(30) PRIMARY KEY,
    account_id VARCHAR(30) NOT NULL REFERENCES accounts(id),
    flag_type VARCHAR(100), -- e.g., 'ineligible', 'high-risk', 'fraud'
    reason TEXT,
    flagged_by INT,
    flagged_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
