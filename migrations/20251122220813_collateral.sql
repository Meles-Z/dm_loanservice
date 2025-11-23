-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE collaterals (
    id VARCHAR(36) PRIMARY KEY,
    account_id VARCHAR(36) NOT NULL REFERENCES accounts(id),
    property_id VARCHAR(36) NOT NULL REFERENCES property(id),
    collateral_type VARCHAR(50) NOT NULL,   -- PROPERTY, CASH, EQUIPMENT
    security_type VARCHAR(50),               -- FIRST_LIEN, SECOND_LIEN
    lien_position VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS collateral;
SELECT 'down SQL query';
-- +goose StatementEnd
