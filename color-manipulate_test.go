package color

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestMix_1 verifies an even mix of black and white, which rounds 127.5 up to 128
func TestMix_1(t *testing.T) {
	color := Parse("#000000").Mix("#ffffff", 50)
	require.Equal(t, "#808080", color.Hex())
}

// TestMix_2 verifies an even mix of blue and white
func TestMix_2(t *testing.T) {
	color := Parse("#0000ff").Mix("#ffffff", 50)
	require.Equal(t, "#8080ff", color.Hex())
}

// TestMix_3 verifies an even mix of blue and black
func TestMix_3(t *testing.T) {
	color := Parse("#0000ff").Mix("#000000", 50)
	require.Equal(t, "#000080", color.Hex())
}

// TestMix_4 verifies a 25% move toward black
func TestMix_4(t *testing.T) {
	color := Parse("#00ffff").Mix("#000000", 25)
	require.Equal(t, "#00bfbf", color.Hex())
}

// TestMix_5 verifies a 75% move toward black
func TestMix_5(t *testing.T) {
	color := Parse("#00ffff").Mix("#000000", 75)
	require.Equal(t, "#004040", color.Hex())
}

// TestMix_Empty verifies that mixing with an empty color returns the other color unchanged
func TestMix_Empty(t *testing.T) {

	blue := Parse("#0000ff")

	require.Equal(t, blue, blue.Mix("not a color", 50))
	require.Equal(t, blue, blue.Mix((*Color)(nil), 50))
	require.Equal(t, blue, NewColor().Mix(blue, 50))
	require.True(t, NewColor().Mix(NewColor(), 50).IsEmpty())
}

// TestMix_Percent verifies the 0% and 100% endpoints, and that other percentages are clamped to them
func TestMix_Percent(t *testing.T) {

	blue := Parse("#0000ff")

	require.Equal(t, "#0000ff", blue.Mix(White, 0).Hex())
	require.Equal(t, "#0000ff", blue.Mix(White, -50).Hex())
	require.Equal(t, "#ffffff", blue.Mix(White, 100).Hex())
	require.Equal(t, "#ffffff", blue.Mix(White, 150).Hex())
}

// TestMix_KeepsOpacity verifies that the result keeps this color's opacity, not the other's
func TestMix_KeepsOpacity(t *testing.T) {

	translucent := Color{Red: 0, Green: 0, Blue: 0, Opacity: 0.5}

	require.Equal(t, Color{Red: 128, Green: 128, Blue: 128, Opacity: 0.5}, translucent.Mix("#ffffff", 50))
}

// TestMix_SameColor verifies that mixing any channel with itself, by any percentage, leaves it unchanged
func TestMix_SameColor(t *testing.T) {

	for value := 0; value <= 255; value++ {
		for percent := 0; percent <= 100; percent++ {
			if result := mix(value, value, percent); result != value {
				require.Equal(t, value, result, "mix(%d, %d, %d)", value, value, percent)
			}
		}
	}
}

// TestMix_Values verifies that Mix accepts the other color in every form that Parse reads
func TestMix_Values(t *testing.T) {

	expected := "#808080"

	require.Equal(t, expected, Black.Mix(White, 50).Hex())
	require.Equal(t, expected, Black.Mix(&White, 50).Hex())
	require.Equal(t, expected, Black.Mix("#ffffff", 50).Hex())
	require.Equal(t, expected, Black.Mix("rgb(255, 255, 255)", 50).Hex())
}

// TestLighten verifies a 50% move from black toward white
func TestLighten(t *testing.T) {
	color := Parse("#000000").Lighten(50)
	require.Equal(t, "#808080", color.Hex())
}

// TestLighten2 verifies a 12% move from black toward white
func TestLighten2(t *testing.T) {
	color := Parse("#000000").Lighten(12)
	require.Equal(t, "#1f1f1f", color.Hex())
}

// TestDarken verifies a 50% move from white toward black
func TestDarken(t *testing.T) {
	color := Parse("#ffffff").Darken(50)
	require.Equal(t, "#808080", color.Hex())
}

// TestDarken2 verifies a 12% move from white toward black
func TestDarken2(t *testing.T) {
	color := Parse("#ffffff").Darken(12)
	require.Equal(t, "#e0e0e0", color.Hex())
}

// TestText verifies black text on light colors, white text on dark ones, and nothing on an empty one
func TestText(t *testing.T) {

	require.Equal(t, Black, gray(255).Text())
	require.Equal(t, Black, gray(128).Text())
	require.Equal(t, White, gray(127).Text())
	require.Equal(t, White, gray(0).Text())
	require.True(t, NewColor().Text().IsEmpty())
}

// TestTextExt verifies every band of TextExt, at both edges of each threshold
func TestTextExt(t *testing.T) {

	test := func(mean int, expected string) {
		t.Helper()
		require.Equal(t, expected, gray(mean).TextExt().Hex(), "mean %d", mean)
	}

	test(255, "#333333")
	test(213, "#2b2b2b")
	test(212, "#151515")
	test(170, "#111111")
	test(169, "#000000")
	test(128, "#000000")
	test(127, "#ffffff")
	test(86, "#ffffff")
	test(85, "#eeeeee")
	test(43, "#eaeaea")
	test(42, "#d4d4d4")
	test(0, "#cccccc")

	require.True(t, NewColor().TextExt().IsEmpty())
}

// TestContrast verifies that light colors darken and dark colors lighten
func TestContrast(t *testing.T) {

	require.Equal(t, "#cccccc", gray(255).Contrast(20).Hex())
	require.Equal(t, "#666666", gray(128).Contrast(20).Hex())
	require.Equal(t, "#999999", gray(127).Contrast(20).Hex())
	require.Equal(t, "#333333", gray(0).Contrast(20).Hex())

	// An empty color counts as dark, so it "lightens" into white
	require.Equal(t, White, NewColor().Contrast(20))
}

// FuzzMix verifies that a mix lands between its two colors, within half a step of the exact
// weighted average, and keeps this color's opacity
func FuzzMix(f *testing.F) {

	f.Add(uint32(0x0000ff), uint32(0xffffff), 50)
	f.Add(uint32(0x030303), uint32(0x030303), 30)
	f.Add(uint32(0x0a141e), uint32(0x28323c), -10)
	f.Add(uint32(0x0a141e), uint32(0x28323c), 1000)

	f.Fuzz(func(t *testing.T, firstRGB uint32, secondRGB uint32, percent int) {

		first := colorFromRGB(firstRGB, 0.5)
		second := colorFromRGB(secondRGB, 1)
		result := first.Mix(second, percent)

		require.Equal(t, first.Opacity, result.Opacity)

		requireMixed(t, result.Red, first.Red, second.Red, percent)
		requireMixed(t, result.Green, first.Green, second.Green, percent)
		requireMixed(t, result.Blue, first.Blue, second.Blue, percent)

		if percent <= 0 {
			require.Equal(t, first, result)
		}

		if percent >= 100 {
			require.Equal(t, second.Hex(), result.Hex())
		}
	})
}

// colorFromRGB returns the color whose channels are the low three bytes of rgb, as in 0xRRGGBB
func colorFromRGB(rgb uint32, opacity float64) Color {
	return Color{
		Red:     int((rgb >> 16) & 0xff),
		Green:   int((rgb >> 8) & 0xff),
		Blue:    int(rgb & 0xff),
		Opacity: opacity,
	}
}

// requireMixed fails the test unless value lies between the two channels, and within half a step
// of the exact weighted average (in hundredths, so the check is exact integer arithmetic)
func requireMixed(t *testing.T, value int, first int, second int, percent int) {
	t.Helper()

	require.GreaterOrEqual(t, value, min(first, second), "channel between %d and %d", first, second)
	require.LessOrEqual(t, value, max(first, second), "channel between %d and %d", first, second)

	percent = max(0, min(100, percent))
	exact := first*(100-percent) + second*percent
	require.LessOrEqual(t, abs(value*100-exact), 50, "channel %d mixing %d and %d by %d%%", value, first, second, percent)
}

// abs returns the absolute value of an integer
func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
