package color

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestIsEmpty verifies that any one channel of -1 makes a color empty
func TestIsEmpty(t *testing.T) {

	test := func(value Color, expected bool) {
		t.Helper()
		require.Equal(t, expected, value.IsEmpty(), "color %#v", value)
		require.Equal(t, !expected, value.NotEmpty(), "color %#v", value)
	}

	test(NewColor(), true)
	test(Color{Red: -1, Green: 0, Blue: 0, Opacity: 1}, true)
	test(Color{Red: 0, Green: -1, Blue: 0, Opacity: 1}, true)
	test(Color{Red: 0, Green: 0, Blue: -1, Opacity: 1}, true)
	test(Black, false)
	test(White, false)

	// Only -1 marks a color empty; other out-of-range values do not
	test(Color{Red: -5, Green: 0, Blue: 0, Opacity: 1}, false)
}

// TestMean verifies that Mean truncates the average channel, and is -1 for an empty color
func TestMean(t *testing.T) {

	require.Equal(t, 2, Color{Red: 1, Green: 2, Blue: 4, Opacity: 1}.Mean())
	require.Equal(t, 0, Black.Mean())
	require.Equal(t, 255, White.Mean())
	require.Equal(t, -1, NewColor().Mean())
}

// TestColorMode verifies the light/dark boundary at a mean of 128, and an empty color's empty mode
func TestColorMode(t *testing.T) {

	test := func(value Color, mode string, isLight bool, isDark bool) {
		t.Helper()
		require.Equal(t, mode, value.ColorMode(), "color %#v", value)
		require.Equal(t, isLight, value.IsLight(), "color %#v", value)
		require.Equal(t, isDark, value.IsDark(), "color %#v", value)
	}

	test(White, ColorModeLight, true, false)
	test(gray(128), ColorModeLight, true, false)
	test(gray(127), ColorModeDark, false, true)
	test(Black, ColorModeDark, false, true)
	test(NewColor(), "", false, false)
}
