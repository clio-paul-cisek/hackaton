package create

// Schedule struct is responsible for mimicking lawpay schedule struct
type Schedule struct {
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
	Schedule    Schedule
}

// RecurringChargeGenerateRequest struct reflect generating request struct
type RecurringChargeGenerateRequest struct {
	Result          string // SUCCESS FAILURE
	RecurringCharge `json:"recurring_charge"`
}
