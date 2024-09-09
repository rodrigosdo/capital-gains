package domain

import (
	"fmt"
)

// TaxRate is the tax rate for capital gains
const TaxRate = 0.20

// Tax represents the tax for a single operation
type Tax struct {
	Tax float64 `json:"tax"`
}

// Taxes is a list of taxes
type Taxes []Tax

// MarshalJSON is a custom JSON marshaller for Tax
// This is a trick hack to round the tax to 2 decimal places
func (t Tax) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"tax":%.2f}`, t.Tax)), nil
}
