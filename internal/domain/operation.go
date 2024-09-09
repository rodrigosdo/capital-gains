package domain

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"
)

// OperationType represents the type of operation.
type OperationType string

// OperationType constants.
const (
	BUY  OperationType = "buy"
	SELL OperationType = "sell"
)

// Operation represents a buy or sell operation.
type Operation struct {
	Quantity int64         `json:"quantity"`
	Type     OperationType `json:"operation"`
	UnitCost float64       `json:"unit-cost"`
}

// Operations represents a list of operations.
type Operations []Operation

// ParseMultilineOperations reads a list of operations from a reader.
// The operations are expected to be in JSON format, one per line.
// The function returns a list of operations or an error if the input is invalid.
// The function stops reading when it finds an empty line.
func ParseMultilineOperations(reader io.Reader) ([]Operations, error) {
	var operations []Operations

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			break
		}

		var o Operations
		if err := json.Unmarshal([]byte(input), &o); err != nil {
			return nil, err
		} else {
			operations = append(operations, o)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return operations, nil
}
