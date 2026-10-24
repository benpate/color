package color

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseRGB_Empty(t *testing.T) {
	result := ParseRGB("")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

func TestParseRGB_Invalid(t *testing.T) {
	result := ParseRGB("123")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

func TestParseRGB_Invalid_Hex(t *testing.T) {
	result := ParseRGB("#123456")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

func TestParseRGB_Invalid_RGB(t *testing.T) {
	result := ParseRGB("rgb(1,2,3,4)")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

func TestParseRGB_Invalid_RGBA(t *testing.T) {
	result := ParseRGB("rgba(1,2,3)")
	require.True(t, result.IsEmpty())
	require.False(t, result.NotEmpty())
	require.Equal(t, -1, result.Red)
	require.Equal(t, -1, result.Green)
	require.Equal(t, -1, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

func TestParseRGB_RGB(t *testing.T) {
	result := ParseRGB("rgb(1,2,3)")
	require.True(t, result.NotEmpty())
	require.False(t, result.IsEmpty())
	require.Equal(t, 1, result.Red)
	require.Equal(t, 2, result.Green)
	require.Equal(t, 3, result.Blue)
	require.Equal(t, 1.0, result.Opacity)
}

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
