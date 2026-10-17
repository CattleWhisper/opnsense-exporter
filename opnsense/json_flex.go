package opnsense

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// flexInt is an integer that tolerates the inconsistent JSON shapes the
// OPNsense MVC API uses for numeric Kea lease fields. Depending on the
// OPNsense release and Kea plugin build, a field such as valid_lifetime,
// expire, state, subnet_id or pool_id may arrive as a bare JSON number
// (86400), a quoted number ("86400"), an empty string ("") or null. All of
// these decode into an int.
//
// Bounds: the token is always parsed at 64-bit width, on every platform,
// so the conversion below can never truncate. That matters on the 32-bit
// release target (linux/arm/v6), where int is 32 bits wide and the
// "infinite" valid_lifetime of 0xffffffff (4294967295) would not fit.
// expire is a unix timestamp. A value that overflows int64, or any
// non-numeric token, is returned as an error rather than silently coerced
// to 0, so a malformed payload surfaces loudly instead of emitting a wrong
// metric. Empty string and null are treated as a legitimate "absent" value
// and decode to 0.
type flexInt int64

func (f *flexInt) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("null")) {
		*f = 0
		return nil
	}

	// Unwrap a quoted value -- "86400" -> 86400 -- decoding the JSON string
	// so any escaping is handled correctly before we parse it.
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return fmt.Errorf("flexInt: %w", err)
		}
		s = strings.TrimSpace(s)
		if s == "" {
			*f = 0
			return nil
		}
		b = []byte(s)
	}

	// Parse at 64-bit width so the conversion below is always lossless.
	// Values that overflow int64 fail here instead.
	n, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return fmt.Errorf("flexInt: cannot parse %q as integer: %w", string(b), err)
	}
	*f = flexInt(n)
	return nil
}

// flexStringSlice is a []string that tolerates the OPNsense MVC API
// returning the Kea is_reserved field either as a JSON array (["hwaddr"])
// or as a bare string ("hwaddr" for a reserved lease, "" otherwise). A
// non-empty string becomes a single-element slice; an empty string or null
// becomes an empty slice, so callers can keep using len() to detect whether
// a lease is reserved.
type flexStringSlice []string

func (s *flexStringSlice) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("null")) {
		*s = nil
		return nil
	}

	if b[0] == '[' {
		var arr []string
		if err := json.Unmarshal(b, &arr); err != nil {
			return fmt.Errorf("flexStringSlice: %w", err)
		}
		*s = arr
		return nil
	}

	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return fmt.Errorf("flexStringSlice: %w", err)
	}
	if strings.TrimSpace(str) == "" {
		*s = nil
		return nil
	}
	*s = flexStringSlice{str}
	return nil
}

// flexCounter is a kernel counter decoded from JSON. A broken FreeBSD
// counter can underflow and wrap to a value close to 2^64, which OPNsense
// then serializes as a float such as 1.8446744073679895e+19. Such a value
// does not fit an int64, and a plain int64 field would fail the whole
// response, dropping every metric of the collector.
//
// flexCounter accepts integers, quoted integers and floats. A value that
// is negative, not finite or does not fit an int64 decodes to 0, so one
// bogus counter only zeroes itself. Empty string and null decode to 0.
type flexCounter int64

func (f *flexCounter) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) > 0 && b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return fmt.Errorf("flexCounter: %w", err)
		}
		b = []byte(strings.TrimSpace(s))
	}
	if len(b) == 0 || bytes.Equal(b, []byte("null")) {
		*f = 0
		return nil
	}

	if n, err := strconv.ParseInt(string(b), 10, 64); err == nil {
		*f = flexCounter(max(n, 0))
		return nil
	}

	v, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		var numErr *strconv.NumError
		if errors.As(err, &numErr) && errors.Is(numErr.Err, strconv.ErrRange) {
			*f = 0
			return nil
		}
		return fmt.Errorf("flexCounter: cannot parse %q as a number: %w", string(b), err)
	}
	// 2^63 is the first float64 above the int64 range.
	if math.IsNaN(v) || v < 0 || v >= math.Exp2(63) {
		*f = 0
		return nil
	}
	*f = flexCounter(v)
	return nil
}
