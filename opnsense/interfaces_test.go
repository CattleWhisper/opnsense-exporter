package opnsense

import "testing"

func TestParseOptionalNumber(t *testing.T) {
	cases := []struct {
		value  string
		suffix string
		want   *float64
	}{
		{"1000000000 bit/s", "bit/s", ptr(1000000000)},
		{"2", "", ptr(2)},
		{" 0 ", "", ptr(0)},
		{"", "", nil},
		{"n/a", "", nil},
	}

	for _, tc := range cases {
		got := parseOptionalNumber(tc.value, tc.suffix)
		switch {
		case tc.want == nil && got != nil:
			t.Errorf("%q: expected nil, got %v", tc.value, *got)
		case tc.want != nil && (got == nil || *got != *tc.want):
			t.Errorf("%q: expected %v, got %v", tc.value, *tc.want, got)
		}
	}
}

func ptr(v float64) *float64 {
	return &v
}
