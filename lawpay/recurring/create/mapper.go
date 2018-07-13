package create

import (
	"encoding/json"
)

// Mapper struct is used to map request to correct response
type Mapper struct {
	recurringChargeGenerateRequest *RecurringChargeGenerateRequest
}

// NewMapper function is used to create new Mapper struct.
func NewMapper(recurringChargeGenerateRequest *RecurringChargeGenerateRequest) *Mapper {
	return &Mapper{
		recurringChargeGenerateRequest: recurringChargeGenerateRequest,
	}
}

// MapRequest maps request to given response
func (m Mapper) MapRequest() ([]byte, error) {
	switch m.recurringChargeGenerateRequest.Result {
	case "SUCCESS":
		return m.success()
	case "FAILURE":
		return m.failure()
	default:
		return m.empty()
	}
}

func (m Mapper) empty() ([]byte, error) {
	var stubData []byte
	return stubData, nil
}

func (m Mapper) success() ([]byte, error) {
	response := newSuccessResponse(m.recurringChargeGenerateRequest.AccountID, m.recurringChargeGenerateRequest.RecurringCharge)
	return json.Marshal(response)
}

func (m Mapper) failure() ([]byte, error) {
	return m.empty()
}
