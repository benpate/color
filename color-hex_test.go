package color

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

func TestParseHex_Valid2(t *testing.T) {
	result := ParseHex("#abcdef")
	require.True(t, result.NotEmpty())
	require.False(t, result.IsEmpty())
	require.Equal(t, 1.0, result.Opacity)
	require.Equal(t, "#abcdef", result.String())
}

func TestFormatHex(t *testing.T) {
	require.Equal(t, "01", formatHex(1))
	require.Equal(t, "0a", formatHex(10))
	require.Equal(t, "ff", formatHex(255))

	require.Equal(t, "", formatHex(-1))
	require.Equal(t, "", formatHex(256))
}
