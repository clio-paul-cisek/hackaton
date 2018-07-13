package create

type schedule struct {
	Start         string
	IntervalUnit  string `json:"interval_unit"`
	IntervalDelay int    `json:"interval_delay"`
}

// RecurringCharge struct reflects request struct
type RecurringCharge struct {
	Description string
	AccountID   string `json:"account_id"`
	Amount      string
	Method      string
	Schedule    schedule
}

// RecurringChargeGenerateRequest struct reflect generating request struct
type RecurringChargeGenerateRequest struct {
	Result          string // SUCCESS FAILURE
	RecurringCharge `json:"recurring_charge"`
}
