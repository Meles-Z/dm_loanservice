-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE property (
    id VARCHAR(36) PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    property_type VARCHAR(50),
    region VARCHAR(50),
    valuation DECIMAL(15,2),
    size_sqft INT,
    year_built INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS property;
SELECT 'down SQL query';
-- +goose StatementEnd
