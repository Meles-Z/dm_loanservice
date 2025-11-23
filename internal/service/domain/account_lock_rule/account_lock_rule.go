package accountlockrule

type AccountLockRule struct {
	ID         string `json:"id"`
	AccountID  string `json:"account_id"`
	Status     string `json:"loan_status"`
	FieldName  string `json:"field_name"`
	LockType   string `json:"lock_type"`
	LockReason string `json:"lock_reason"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
