package create

type schedule struct {
	Start         string
	IntervalUnit  string `json:"interval_unit"`
	IntervalDelay int    `json:"interval_delay"`
}

type RecurringCharge struct {
	Description string
	AccountID   string `json:"account_id"`
	Amount      string
	Method      string
	Schedule    schedule
}

type RecurringChargeGenerateRequest struct {
	Result          string // SUCCESS FAILURE
	RecurringCharge `json:"recurring_charge"`
}