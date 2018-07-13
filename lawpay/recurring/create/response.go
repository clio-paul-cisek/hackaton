package create

// Method struct reflects payment method structure
type Method struct {
	Type        string
	Number      string
	Fingerprint string
	CardType    string `json:"card_type"`
	ExpMonth    int    `json:"exp_month"`
	ExpYear     int    `json:"exp_year"`
	Name        string
}

// Occurrence struct reflects occurrence structure
type Occurrence struct {
	ID       string
	Amount   int
	Status   string
	DueDate  string `json:"due_date"`
	Attempts int
}

// SuccessResponse reflects response on success
type SuccessResponse struct {
	ID               string
	Status           string
	AccountID        string `json:"account_id"`
	Method           Method
	Schedule         Schedule
	Description      string
	Amount           int
	Currency         string
	TotalOccurrences int    `json:"total_occurrences"`
	TotalAmount      int    `json:"total_amount"`
	NextPayment      string `json:"next_payment"`
	Occurrences      []Occurrence
}
