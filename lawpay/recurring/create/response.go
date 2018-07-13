package create

import (
	"strconv"
)

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
	ID        string
	Status    string
	AccountID string `json:"account_id"`
	Method
	Schedule
	Description      string
	Amount           int
	Currency         string
	TotalOccurrences int    `json:"total_occurrences"`
	TotalAmount      int    `json:"total_amount"`
	NextPayment      string `json:"next_payment"`
	Occurrences      []Occurrence
}

func newSuccessResponse(UID string, rc RecurringCharge) *SuccessResponse {
	amount, _ := strconv.Atoi(rc.Amount)
	return &SuccessResponse{
		ID:        UID,
		Status:    "Active",
		AccountID: rc.AccountID,
		// THIS should be provided by enpoint
		Method: Method{
			Type:        "card",
			Number:      "4242424242424242",
			Fingerprint: "GunPelYVthifNV63LEw1",
			CardType:    "VISA",
			ExpMonth:    10,
			ExpYear:     2022,
			Name:        "Test Customer",
		},
		Schedule:         rc.Schedule,
		Description:      rc.Description,
		Amount:           amount,
		Currency:         "USD",
		TotalOccurrences: 0,
		TotalAmount:      0,
		NextPayment:      "2016-01-01",
		Occurrences: []Occurrence{
			Occurrence{
				ID:       "_LIG1tsDQZ21oBgPYTRJdQ",
				Amount:   1250,
				Status:   "PENDING",
				DueDate:  "2016-01-01",
				Attempts: 0,
			},
		},
	}
}
