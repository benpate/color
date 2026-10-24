package color

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// gray returns a fully opaque gray whose channels, and therefore its Mean, all equal value
func gray(value int) Color {
	return Color{Red: value, Green: value, Blue: value, Opacity: 1}
}

// requireChannel fails the test unless value is -1 or a valid 0-255 channel
func requireChannel(t *testing.T, value int) {
	t.Helper()
	require.True(t, (value == -1) || isValidInt(value), "channel %d", value)
}

// TestNewColor verifies that a new Color is empty and fully opaque
func TestNewColor(t *testing.T) {

	result := NewColor()

	require.Equal(t, Color{Red: -1, Green: -1, Blue: -1, Opacity: 1}, result)
	require.True(t, result.IsEmpty())
}

// TestParse verifies every kind of value that Parse accepts, and what it does with the rest
func TestParse(t *testing.T) {

	original := Color{Red: 1, Green: 2, Blue: 3, Opacity: 0.5}

	test := func(value any, expected Color) {
		t.Helper()
		require.Equal(t, expected, Parse(value), "value %#v", value)
	}

	test(original, original)
	test(&original, original)
	test("#010203", Color{Red: 1, Green: 2, Blue: 3, Opacity: 1})
	test("010203", Color{Red: 1, Green: 2, Blue: 3, Opacity: 1})
	test("rgb(4,5,6)", Color{Red: 4, Green: 5, Blue: 6, Opacity: 1})
	test("rgba(4, 5, 6, 0.25)", Color{Red: 4, Green: 5, Blue: 6, Opacity: 0.25})

	// A partly readable rgb() string keeps the channels it could read
	test("rgb(1,x,3)", Color{Red: 1, Green: -1, Blue: 3, Opacity: 1})

	test("not a color", NewColor())
	test("", NewColor())
	test(5, NewColor())
	test(nil, NewColor())
	test([]byte("#010203"), NewColor())
}

// TestParse_NilPointer verifies that a nil *Color is an empty Color rather than a panic
func TestParse_NilPointer(t *testing.T) {

	var value *Color

	require.NotPanics(t, func() { Parse(value) })
	require.Equal(t, NewColor(), Parse(value))
}

// FuzzParse verifies that Parse reads a string as hex when it can, and as rgb() otherwise
func FuzzParse(f *testing.F) {

	f.Add("#010203")
	f.Add("rgb(1,2,3)")
	f.Add("rgba(1, 2, 3, 0.5)")
	f.Add("")
	f.Add("#+f+f+f")

	f.Fuzz(func(t *testing.T, value string) {

		result := Parse(value)

		if hex := ParseHex(value); hex.NotEmpty() {
			require.Equal(t, hex, result)
			return
		}

		require.Equal(t, ParseRGB(value), result)
	})
}
