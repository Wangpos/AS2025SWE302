package shipping

import (
	"strings"
	"testing"
)

// Equivalence Partitioning Test for CalculateShippingFeeV2
func TestCalculateShippingFeeV2_EquivalencePartitioning(t *testing.T) {
	// P1: Too Small (Invalid)
	_, err := CalculateShippingFeeV2(-5, "Domestic", false)
	if err == nil {
		t.Error("Test failed: Expected an error for a negative weight, but got nil")
	}

	// P2 & P4: Valid weight and valid zone
	fee, err := CalculateShippingFeeV2(10, "Domestic", false)
	if err != nil {
		t.Errorf("Test failed: Expected no error for a valid weight, but got %v", err)
	}
	expectedFee := 5.0 // Domestic base fee, standard package, not insured
	if fee != expectedFee {
		t.Errorf("Test failed: Expected fee of %f, but got %f", expectedFee, fee)
	}

	// P3: Too Large (Invalid)
	_, err = CalculateShippingFeeV2(100, "International", false)
	if err == nil {
		t.Error("Test failed: Expected an error for an overweight package, but got nil")
	}

	// P5: Invalid zone
	_, err = CalculateShippingFeeV2(20, "Local", false)
	if err == nil {
		t.Error("Test failed: Expected an error for an invalid zone, but got nil")
	}
}

// Boundary Value Analysis Test for CalculateShippingFeeV2
func TestCalculateShippingFeeV2_BoundaryValueAnalysis(t *testing.T) {
	testCases := []struct {
		name        string
		weight      float64
		zone        string
		expectError bool
	}{
		{"Weight at lower invalid boundary", 0, "Domestic", true},
		{"Weight just above lower boundary", 0.1, "Domestic", false},
		{"Weight at upper valid boundary", 50, "International", false},
		{"Weight just above upper boundary", 50.1, "Express", true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := CalculateShippingFeeV2(tc.weight, tc.zone, false)
			if tc.expectError && err == nil {
				t.Errorf("Expected an error, but got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
		})
	}
}

// Decision Table Testing for CalculateShippingFeeV2
func TestCalculateShippingFeeV2_DecisionTable(t *testing.T) {
	testCases := []struct {
		name          string
		weight        float64
		zone          string
		insured       bool
		expectedFee   float64
		expectedError string // substring to check in error
	}{
		// Rule 1: Invalid weight. Zone does not matter.
		{"Rule 1: Weight too low", -10, "Domestic", false, 0, "invalid weight"},
		{"Rule 1: Weight too high", 60, "International", false, 0, "invalid weight"},

		// Rule 2: Valid weight, Domestic zone
		{"Rule 2: Domestic", 10, "Domestic", false, 5.0, ""},
		{"Rule 2: Domestic Insured", 10, "Domestic", true, 5.0 * 1.015, ""},

		// Rule 3: Valid weight, International zone
		{"Rule 3: International", 10, "International", false, 20.0, ""},
		{"Rule 3: International Insured", 10, "International", true, 20.0 * 1.015, ""},

		// Rule 4: Valid weight, Express zone
		{"Rule 4: Express", 10, "Express", false, 30.0, ""},
		{"Rule 4: Express Insured", 10, "Express", true, 30.0 * 1.015, ""},

		// Rule 5: Valid weight, Invalid Zone
		{"Rule 5: Invalid Zone", 10, "Unknown", false, 0, "invalid zone"},
		{"Rule 5: Invalid Zone Insured", 10, "Unknown", true, 0, "invalid zone"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fee, err := CalculateShippingFeeV2(tc.weight, tc.zone, tc.insured)
			if tc.expectedError != "" {
				if err == nil {
					t.Fatalf("Expected error containing '%s', but got nil", tc.expectedError)
				}
				if !strings.Contains(err.Error(), tc.expectedError) {
					t.Errorf("Expected error to contain '%s', got '%s'", tc.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("Expected no error, but got: %v", err)
				}
				if fee != tc.expectedFee {
					t.Errorf("Expected fee %f, but got %f", tc.expectedFee, fee)
				}
			}
		})
	}
}

