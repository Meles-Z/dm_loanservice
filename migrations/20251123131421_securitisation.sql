-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE securitisation_pools (
    id VARCHAR(36) PRIMARY KEY,
    funding_source VARCHAR NOT NULL,
    servicing_role VARCHAR NOT NULL,
    spv_name VARCHAR(255) NOT NULL,
    spv_jurisdiction VARCHAR NOT NULL,
    pool_allocation_date DATE NOT NULL,
    loan_transfer_date DATE NOT NULL,
    current_pool_balance NUMERIC(18, 2) NOT NULL DEFAULT 0.00,
    factor NUMERIC(10, 6) NOT NULL CHECK (factor >= 0 AND factor <= 1),
    note_class VARCHAR(10) NOT NULL,
    interest_remittance_date DATE NOT NULL,
    principal_remittance_date DATE NOT NULL,
    servicing_fee_rate NUMERIC(6, 4) NOT NULL CHECK (servicing_fee_rate >= 0),
    reporting_currency VARCHAR(3) NOT NULL,
    esma_asset_code VARCHAR(50),
    credit_enhancement_type VARCHAR NOT NULL,
    investor_report_identifier VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
