package internal_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"capital-gains/cmd/cli/internal"
)

func TestRun(t *testing.T) {
	tests := []struct {
		inputFile    string
		expectedFile string
	}{
		{"testdata/case_1_input.golden", "testdata/case_1_output.golden"},
		{"testdata/case_2_input.golden", "testdata/case_2_output.golden"},
		{"testdata/case_1+2_input.golden", "testdata/case_1+2_output.golden"},
		{"testdata/case_3_input.golden", "testdata/case_3_output.golden"},
		{"testdata/case_4_input.golden", "testdata/case_4_output.golden"},
		{"testdata/case_5_input.golden", "testdata/case_5_output.golden"},
		{"testdata/case_6_input.golden", "testdata/case_6_output.golden"},
		{"testdata/case_7_input.golden", "testdata/case_7_output.golden"},
		{"testdata/case_8_input.golden", "testdata/case_8_output.golden"},
	}

	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			t.Parallel()

			input, err := os.Open(tt.inputFile)
			if err != nil {
				t.Fatalf("failed to open input file: %v", err)
			}
			defer func(input *os.File) {
				err := input.Close()
				if err != nil {
					t.Fatalf("failed to close input file: %v", err)
				}
			}(input)

			expected, err := os.ReadFile(tt.expectedFile)
			if err != nil {
				t.Fatalf("failed to read expected file: %v", err)
			}

			var output bytes.Buffer
			if err := internal.Run(input, &output); err != nil {
				t.Fatalf("failed to run: %v", err)
			}

			assert.Equal(t, string(expected), output.String())
		})
	}
}
