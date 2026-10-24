package color

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParseHex_Empty verifies that an empty string is an empty color
func TestParseHex_Empty(t *testing.T) {
	result := ParseHex("")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
	require.Equal(t, "", result.String())
}

// TestParseHex_Invalid verifies that a value of the wrong length is an empty color
func TestParseHex_Invalid(t *testing.T) {
	result := ParseHex("123")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
	require.Equal(t, "", result.String())
}

// TestParseHex_RGBA verifies that ParseHex does not read rgba() strings
func TestParseHex_RGBA(t *testing.T) {
	result := ParseHex("rgba(1,2,3,0)")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
	require.Equal(t, "", result.String())
}

// TestParseHex_Valid verifies a six-digit color with its leading "#"
func TestParseHex_Valid(t *testing.T) {
	result := ParseHex("#123456")
	require.True(t, result.NotEmpty())
	require.False(t, result.IsEmpty())
	require.Equal(t, 0x12, result.Red)
	require.Equal(t, 0x34, result.Green)
	require.Equal(t, 0x56, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
	require.Equal(t, "#123456", result.String())
}

// TestParseHex_Valid2 verifies that lowercase letter digits parse and format back unchanged
func TestParseHex_Valid2(t *testing.T) {
	result := ParseHex("#abcdef")
	require.True(t, result.NotEmpty())
	require.False(t, result.IsEmpty())
	require.Equal(t, 1.0, result.Opacity)
	require.Equal(t, "#abcdef", result.String())
}

// TestParseHex_NoHash verifies that the leading "#" is optional
func TestParseHex_NoHash(t *testing.T) {
	require.Equal(t, Color{Red: 0x12, Green: 0x34, Blue: 0x56, Opacity: 1}, ParseHex("123456"))
}

// TestParseHex_Uppercase verifies that uppercase digits parse, and format back as lowercase
func TestParseHex_Uppercase(t *testing.T) {
	result := ParseHex("#ABCDEF")
	require.Equal(t, Color{Red: 0xab, Green: 0xcd, Blue: 0xef, Opacity: 1}, result)
	require.Equal(t, "#abcdef", result.String())
}

// TestParseHex_PartlyInvalid verifies that one bad digit pair empties the color but keeps the others
func TestParseHex_PartlyInvalid(t *testing.T) {
	result := ParseHex("#12zz56")
	require.Equal(t, Color{Red: 0x12, Green: -1, Blue: 0x56, Opacity: 1}, result)
	require.True(t, result.IsEmpty())
}

// TestParseHex_Unsupported verifies forms ParseHex does not read: CSS shorthand, alpha, and spacing
func TestParseHex_Unsupported(t *testing.T) {

	test := func(value string) {
		t.Helper()
		require.True(t, ParseHex(value).IsEmpty(), "value %q", value)
	}

	test("#fff")
	test("#12345678")
	test(" #123456")
	test("#123456 ")
	test("##123456")
	test("#ééé")
}

// TestParseHex_Signs pins strconv's sign handling: a "+" digit pair parses, and a "-" pair does not
func TestParseHex_Signs(t *testing.T) {
	require.Equal(t, Color{Red: 15, Green: 15, Blue: 15, Opacity: 1}, ParseHex("#+f+f+f"))
	require.True(t, ParseHex("#-1-1-1").IsEmpty())
}

// TestFormatHex verifies two-digit padding and the empty result outside 0-255
func TestFormatHex(t *testing.T) {
	require.Equal(t, "01", formatHex(1))
	require.Equal(t, "0a", formatHex(10))
	require.Equal(t, "ff", formatHex(255))

	require.Equal(t, "", formatHex(-1))
	require.Equal(t, "", formatHex(256))
}

// FuzzParseHex verifies that every channel is -1 or 0-255, and that a parsed color survives Hex
func FuzzParseHex(f *testing.F) {

	f.Add("#123456")
	f.Add("abcdef")
	f.Add("#12zz56")
	f.Add("#+f+f+f")
	f.Add("")
	f.Add("#ééé")

	f.Fuzz(func(t *testing.T, value string) {

		result := ParseHex(value)

		requireChannel(t, result.Red)
		requireChannel(t, result.Green)
		requireChannel(t, result.Blue)
		require.Equal(t, 1.0, result.Opacity)

		if result.IsEmpty() {
			return
		}

		require.Len(t, result.Hex(), 7)
		require.Equal(t, result, ParseHex(result.Hex()))
	})
}
