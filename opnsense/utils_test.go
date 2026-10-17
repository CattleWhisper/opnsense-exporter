package opnsense

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/prometheus/common/promslog"
)

func TestParsePercentage(t *testing.T) {
	logger := promslog.NewNopLogger()
	testRegex := regexp.MustCompile(`\d\.\d %`)

	tests := []struct {
		name           string
		value          string
		regex          *regexp.Regexp
		replacePattern string
		valueTypeName  string
		gatewayName    string
		expected       float64
	}{
		{
			name:           "Valid percentage with space",
			value:          "50.5 %",
			regex:          testRegex,
			replacePattern: " %",
			valueTypeName:  "loss",
			gatewayName:    "Gateway1",
			expected:       50.5,
		},
		{
			name:           "Valid percentage with space",
			value:          "5.5 %",
			regex:          testRegex,
			replacePattern: " %",
			valueTypeName:  "loss",
			gatewayName:    "Gateway1",
			expected:       5.5,
		},
		{
			name:           "Invalid percentage format",
			value:          "invalid %",
			regex:          testRegex,
			replacePattern: " %",
			valueTypeName:  "loss",
			gatewayName:    "Gateway1",
			expected:       -1.0,
		},
		{
			name:           "Invalid regex match (no space)",
			value:          "50.5%",
			regex:          testRegex,
			replacePattern: " %",
			valueTypeName:  "loss",
			gatewayName:    "Gateway1",
			expected:       -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := parseStringToFloatWithReplace(tc.value, tc.regex, tc.replacePattern, tc.valueTypeName, logger)
			if result != tc.expected {
				t.Errorf("parsePercentage(%s, %v, %s, %s, logger, %s) = %v; want %v",
					tc.value, tc.regex, tc.replacePattern, tc.valueTypeName, tc.gatewayName, result, tc.expected)
			}
		})
	}
}

func TestSliceIntToMapStringInt(t *testing.T) {
	input := []string{"1", "2", "3"}
	expected := map[string]int64{"1": 1, "2": 2, "3": 3}

	result, _ := sliceIntToMapStringInt(input, "test")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestParseStringToBool(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{
			name:     "Zero",
			value:    "0",
			expected: false,
		},
		{
			name:     "One",
			value:    "1",
			expected: true,
		},
		{
			name:     "Invalid/Unknown",
			value:    "2",
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := parseStringToBool(tc.value)
			if result != tc.expected {
				t.Errorf("parseStringToBool(%s) = %v; want %v",
					tc.value, result, tc.expected)
			}
		})
	}
}

// Counters above 2^31 must parse on 32-bit platforms (linux/arm/v6) too.
func TestParseStringToIntLargeCounter(t *testing.T) {
	got, err := parseStringToInt("123456789012", "test")
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}
	if got != 123456789012 {
		t.Errorf("expected 123456789012, got %d", got)
	}
}
