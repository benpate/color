package color

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParseFloat verifies the 0-1 range, trimming, and every kind of rejection
func TestParseFloat(t *testing.T) {

	test := func(value string, expected float64) {
		t.Helper()
		require.Equal(t, expected, parseFloat(value), "value %q", value)
	}

	test("0", 0)
	test("1", 1)
	test("0.5", 0.5)
	test(" 0.5 ", 0.5)
	test("1.01", -1)
	test("-0.1", -1)
	test("NaN", -1)
	test("Inf", -1)
	test("-Inf", -1)
	test("x", -1)
	test("", -1)
}

// TestParseFloat_NegativeZero pins that "-0" passes the range check as negative zero
func TestParseFloat_NegativeZero(t *testing.T) {

	result := parseFloat("-0")

	require.Zero(t, result)
	require.True(t, math.Signbit(result))
}

// TestIsValidInt verifies the inclusive 0-255 channel range
func TestIsValidInt(t *testing.T) {

	require.False(t, isValidInt(-1))
	require.True(t, isValidInt(0))
	require.True(t, isValidInt(255))
	require.False(t, isValidInt(256))
}
