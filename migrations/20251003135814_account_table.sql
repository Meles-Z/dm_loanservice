-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TYPE forbearance_type AS ENUM ('Payment Holiday', 'Interest Only', 'Reduced Payment');
CREATE TYPE collateral_type AS ENUM ('House', 'Flat', 'Office', 'Land');
CREATE TYPE security_type AS ENUM ('Mortgage Deed', 'Legal Charge', 'Guarantee');
CREATE TYPE stress_test_result AS ENUM ('Pass', 'Fail');


CREATE TABLE accounts (
    id VARCHAR(36) PRIMARY KEY,
    customer_id VARCHAR(36) NOT NULL REFERENCES customers(id),
    product_id VARCHAR(36) NOT NULL REFERENCES products(id),
    loan_amount DECIMAL(15,2) NOT NULL,
    balance_outstanding DECIMAL(15,2) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    term_years INT NOT NULL,

    arrears_flag BOOLEAN DEFAULT FALSE,
    arrears_amount DECIMAL(15,2),
    arrears_days INT,

    forbearance_flag BOOLEAN DEFAULT FALSE,
    forbearance_type forbearance_type,

    fraud_flag BOOLEAN DEFAULT FALSE,
    fraud_notes TEXT,

    redraw_facility BOOLEAN DEFAULT FALSE,
    collateral_address VARCHAR(255),
    collateral_type collateral_type,
    security_type security_type,

    portfolio_id VARCHAR(36),
    stress_test_result stress_test_result,
    capital_adequacy_flag BOOLEAN DEFAULT FALSE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);


-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS forbearance_type CASCADE;
DROP TYPE IF EXISTS collateral_type CASCADE;
DROP TYPE IF EXISTS security_type CASCADE;
DROP TYPE IF EXISTS stress_test_result CASCADE;
DROP TABLE IF EXISTS accounts;
SELECT 'down SQL query';
-- +goose StatementEnd




/*
| Field                   | Type                                                         | Explanation                                                                                                                                                                                        |
| ----------------------- | ------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `id`                    | VARCHAR(36)                                                  | Unique identifier for the account (loan). Usually a UUID.                                                                                                                                          |
| `customer_id`           | VARCHAR(36)                                                  | Links the loan to a customer in the `customers` table.                                                                                                                                             |
| `product_id`            | VARCHAR(36)                                                  | Links the loan to a specific product (e.g., mortgage, personal loan). Used for rules like interest rate, frequency, etc.                                                                           |
| `loan_amount`           | DECIMAL(15,2)                                                | The total original amount of the loan borrowed.                                                                                                                                                    |
| `balance_outstanding`   | DECIMAL(15,2)                                                | How much the borrower still owes. Decreases as they make repayments.                                                                                                                               |
| `start_date`            | DATE                                                         | Date the loan officially started (first day of the loan).                                                                                                                                          |
| `end_date`              | DATE                                                         | Date the loan is scheduled to end (maturity date).                                                                                                                                                 |
| `term_years`            | INT                                                          | Duration of the loan in years.                                                                                                                                                                     |
| `arrears_flag`          | BOOLEAN                                                      | TRUE if the borrower has missed payments or is behind.                                                                                                                                             |
| `arrears_amount`        | DECIMAL(15,2)                                                | Total money overdue if the borrower is in arrears.                                                                                                                                                 |
| `arrears_days`          | INT                                                          | Number of days the borrower is late in payments.                                                                                                                                                   |
| `forbearance_flag`      | BOOLEAN, halaa rakkissaa kessaa yoo jiratee akkati kafkatuu                                                      | TRUE if the borrower is under a special hardship program.                                                                                                                                          |
| `forbearance_type`      | ENUM ('Payment Holiday', 'Interest Only', 'Reduced Payment') | Type of hardship support: <br>• **Payment Holiday** → skip payments temporarily <br>• **Interest Only** → pay only interest, not principal <br>• **Reduced Payment** → pay less than normal amount |
| `fraud_flag`            | BOOLEAN                                                      | TRUE if this account is under fraud investigation.                                                                                                                                                 |
| `fraud_notes`           | TEXT                                                         | Details about fraud investigation or alerts.                                                                                                                                                       |
| `redraw_facility`       | BOOLEAN                                                      | TRUE if borrower can withdraw extra funds from the loan (common in mortgages).                                                                                                                     |
| `collateral_address`    | VARCHAR(255)                                                 | Address of the collateral/property securing the loan.                                                                                                                                              |
| `collateral_type`       | ENUM ('House', 'Flat', 'Office', 'Land')                     | Type of collateral securing the loan.                                                                                                                                                              |
| `security_type`         | ENUM ('Mortgage Deed', 'Legal Charge', 'Guarantee')          | Legal type of security for the loan.                                                                                                                                                               |
| `portfolio_id`          | VARCHAR(36)                                                  | Links the account to a portfolio or business unit (used for reporting).                                                                                                                            |
| `stress_test_result`    | ENUM('Pass', 'Fail')                                         | Indicates whether the account passes the bank’s stress test (financial stability check).                                                                                                           |
| `capital_adequacy_flag` | BOOLEAN                                                      | TRUE if the loan affects capital adequacy reporting. Banks must check capital requirements.                                                                                                        |
| `created_at`            | TIMESTAMP                                                    | When this account was created in the system.                                                                                                                                                       |
| `updated_at`            | TIMESTAMP                                                    | Last time the account details were updated.                                                                                                                                                        |
| `deleted_at`            | TIMESTAMP                                                    | If set, the account is considered deleted/closed. Otherwise NULL.                                                                                                                                  |

*/