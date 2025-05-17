package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// This unit test file was generated with JetBrains AI
func TestComparePerformance(t *testing.T) {
	// Define test cases with actual sizes to test
	tests := []struct {
		name string
		size int
	}{
		{"smallSize", 1000},
		{"mediumSize", 10000},
		{"largeSize", 100000},
		{"extraLargeSize", 1000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout to verify output
			oldStdout := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Failed to create pipe: %v", err)
			}
			os.Stdout = w

			// Call the function that handles a specific size
			comparePerformance(tt.size)

			// Restore stdout and get output
			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			_, err = io.Copy(&buf, r)
			if err != nil {
				t.Fatalf("Failed to read from pipe: %v", err)
			}
			output := buf.String()

			// Verify the expected size is in the output
			expectedSizeText := fmt.Sprintf("=== Performance Comparison for %d nodes ===", tt.size)
			if !strings.Contains(output, expectedSizeText) {
				t.Errorf("Expected output to contain %q for size %d, but got: %s",
					expectedSizeText, tt.size, output)
			}

			// Verify the results match
			if !strings.Contains(output, "Results match") {
				t.Errorf("Results don't match for size %d", tt.size)
			}
		})
	}
}
