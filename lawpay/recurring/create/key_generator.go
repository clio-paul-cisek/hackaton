package create

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

// KeyGenerator struct is used to contain values used for key generation
type KeyGenerator struct {
	AccountID string `json:"account_id"`
	Schedule
	Description string
}

// NewKeyGenerator function creates new KeyGenerator struct
func NewKeyGenerator(accountID, description, start, intervalUnit string, intervalDelay int) KeyGenerator {
	return KeyGenerator{
		AccountID:   accountID,
		Description: description,
		Schedule: Schedule{
			Start:         start,
			IntervalUnit:  intervalUnit,
			IntervalDelay: intervalDelay,
		},
	}
}

// GenerateKey function return hash of KayGenerator struct
func (kg KeyGenerator) GenerateKey() string {
	record := kg.recordToHash()
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (kg KeyGenerator) recordToHash() string {
	return kg.AccountID +
		kg.IntervalUnit +
		strconv.Itoa(kg.IntervalDelay) +
		kg.Start +
		kg.Description
}
