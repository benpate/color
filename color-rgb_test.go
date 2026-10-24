package color

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParseRGB_Empty verifies that an empty string is an empty color
func TestParseRGB_Empty(t *testing.T) {
	result := ParseRGB("")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

// TestParseRGB_Invalid verifies that a value without "rgb(" or "rgba(" is an empty color
func TestParseRGB_Invalid(t *testing.T) {
	result := ParseRGB("123")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

// TestParseRGB_Invalid_Hex verifies that ParseRGB does not read hex strings
func TestParseRGB_Invalid_Hex(t *testing.T) {
	result := ParseRGB("#123456")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

// TestParseRGB_Invalid_RGB verifies that rgb() with four values is an empty color
func TestParseRGB_Invalid_RGB(t *testing.T) {
	result := ParseRGB("rgb(1,2,3,4)")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

// TestParseRGB_Invalid_RGBA verifies that rgba() with three values is an empty color
func TestParseRGB_Invalid_RGBA(t *testing.T) {
	result := ParseRGB("rgba(1,2,3)")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

// TestParseRGB_RGB verifies an rgb() color, which is fully opaque
func TestParseRGB_RGB(t *testing.T) {
	result := ParseRGB("rgb(1,2,3)")
	require.True(t, result.NotEmpty())
	require.False(t, result.IsEmpty())
	require.Equal(t, 1, result.Red)
	require.Equal(t, 2, result.Green)
	require.Equal(t, 3, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

// TestParseRGB_RGBA verifies an rgba() color and its opacity
func TestParseRGB_RGBA(t *testing.T) {
	result := ParseRGB("rgba(1,2,3,0.5)")
	require.True(t, result.NotEmpty())
	require.False(t, result.IsEmpty())
	require.Equal(t, 1, result.Red)
	require.Equal(t, 2, result.Green)
	require.Equal(t, 3, result.Blue)
	require.Equal(t, 0.5, result.Opacity)
}

// TestParseRGB_HighChannels pins the full 0-255 range.  A signed 8-bit
// ParseInt silently failed every channel above 127, so bright colors
// parsed as empty.
func TestParseRGB_HighChannels(t *testing.T) {

	tests := []struct {
		name  string
		value string
		red   int
		green int
		blue  int
	}{
		{"white", "rgb(255,255,255)", 255, 255, 255},
		{"just above the signed-8-bit ceiling", "rgb(128,128,128)", 128, 128, 128},
		{"at the signed-8-bit ceiling", "rgb(127,127,127)", 127, 127, 127},
		{"mixed high and low", "rgb(200,10,90)", 200, 10, 90},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ParseRGB(test.value)
			require.True(t, result.NotEmpty())
			require.Equal(t, test.red, result.Red)
			require.Equal(t, test.green, result.Green)
			require.Equal(t, test.blue, result.Blue)
		})
	}
}

// TestParseRGB_Spaces pins the standard CSS spelling, which puts a space
// after each comma.
func TestParseRGB_Spaces(t *testing.T) {

	result := ParseRGB("rgb(255, 128, 0)")
	require.True(t, result.NotEmpty())
	require.Equal(t, 255, result.Red)
	require.Equal(t, 128, result.Green)
	require.Equal(t, 0, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

// TestParseRGBA_Spaces pins the spaced CSS spelling for the alpha channel too.
func TestParseRGBA_Spaces(t *testing.T) {

	result := ParseRGB("rgba(255, 128, 0, 0.5)")
	require.True(t, result.NotEmpty())
	require.Equal(t, 255, result.Red)
	require.Equal(t, 128, result.Green)
	require.Equal(t, 0, result.Blue)
	require.Equal(t, 0.5, result.Opacity)
}

// TestParseRGB_OutOfRange confirms that range enforcement still rejects
// values above 255, now that ParseInt is no longer doing it by accident.
func TestParseRGB_OutOfRange(t *testing.T) {

	for _, value := range []string{"rgb(256,0,0)", "rgb(999,0,0)", "rgb(-1,0,0)"} {
		result := ParseRGB(value)
		require.Equal(t, -1, result.Red, value)
	}
}

// TestParseRGB_PartlyInvalid verifies that one bad channel empties the color but keeps the others
func TestParseRGB_PartlyInvalid(t *testing.T) {
	result := ParseRGB("rgb(1,x,3)")
	require.Equal(t, Color{Red: 1, Green: -1, Blue: 3, Opacity: 1}, result)
	require.True(t, result.IsEmpty())
}

// TestParseRGBA_InvalidOpacity pins that a bad opacity becomes -1 while the color stays non-empty,
// so String writes it as an rgba() value with an alpha of -1
func TestParseRGBA_InvalidOpacity(t *testing.T) {

	test := func(value string) {
		t.Helper()
		result := ParseRGB(value)
		require.Equal(t, Color{Red: 1, Green: 2, Blue: 3, Opacity: -1}, result, "value %q", value)
		require.True(t, result.NotEmpty(), "value %q", value)
		require.Equal(t, "rgba(1,2,3,-1)", result.String(), "value %q", value)
	}

	test("rgba(1,2,3,x)")
	test("rgba(1,2,3,2)")
	test("rgba(1,2,3,50%)")
}

// TestParseRGB_Unsupported verifies forms ParseRGB does not read: other spellings, and stray spacing
func TestParseRGB_Unsupported(t *testing.T) {

	test := func(value string) {
		t.Helper()
		require.True(t, ParseRGB(value).IsEmpty(), "value %q", value)
	}

	test("RGB(1,2,3)")
	test(" rgb(1,2,3)")
	test("rgb(1 2 3)")
	test("rgb(10%,20%,30%)")

	// A space after the closing parenthesis keeps it on the last value, which then fails
	test("rgb(1,2,3) ")
}

// TestParseInt verifies the 0-255 range, trimming, and what strconv accepts and rejects
func TestParseInt(t *testing.T) {

	test := func(value string, expected int) {
		t.Helper()
		require.Equal(t, expected, parseInt(value), "value %q", value)
	}

	test("0", 0)
	test("255", 255)
	test(" 7 ", 7)
	test("08", 8)
	test("+5", 5)
	test("256", -1)
	test("-1", -1)
	test("32768", -1)
	test("0x10", -1)
	test("1e2", -1)
	test("1.5", -1)
	test("", -1)
}

// FuzzParseRGB verifies that every channel is -1 or 0-255, that opacity is -1 or 0-1,
// and that a parsed color survives RGBA
func FuzzParseRGB(f *testing.F) {

	f.Add("rgb(1,2,3)")
	f.Add("rgba(255, 128, 0, 0.5)")
	f.Add("rgba(1,2,3,x)")
	f.Add("rgb(1,x,3)")
	f.Add("rgba(0,0,0,-0)")
	f.Add("")

	f.Fuzz(func(t *testing.T, value string) {

		result := ParseRGB(value)

		requireChannel(t, result.Red)
		requireChannel(t, result.Green)
		requireChannel(t, result.Blue)
		require.True(t, (result.Opacity == -1) || ((result.Opacity >= 0) && (result.Opacity <= 1)), "opacity %v", result.Opacity)

		if result.IsEmpty() {
			return
		}

		if result.Opacity == -1 {
			return
		}

		require.Equal(t, result, ParseRGB(result.RGBA()))
	})
}
