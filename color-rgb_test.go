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
