package internal

import (
	"encoding/json"
	"fmt"
	"io"

	"capital-gains/internal/domain"
)

// Run reads operations from the reader and writes the calculated taxes to the writer.
// Each line of the reader should contain a JSON array of operations.
func Run(reader io.Reader, writer io.Writer) error {
	multilineOperations, err := domain.ParseMultilineOperations(reader)
	if err != nil {
		return err
	}

	for _, operations := range multilineOperations {
		portfolio := domain.Portfolio{}
		taxes := portfolio.CalculateTaxes(operations)

		encodedTaxes, err := json.Marshal(taxes)
		if err != nil {
			return err
		}

		if _, err := fmt.Fprintln(writer, string(encodedTaxes)); err != nil {
			return err
		}
	}

	return nil
}
