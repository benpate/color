package color

/******************************************
 * Color Manipulations
 ******************************************/

// Mix combines this color with another one, moving towards
// the 'other' color by a specified percentage.  0% will return
// this color, 100% will return the 'other' color.
func (c Color) Mix(other any, percent int) Color {

	otherColor := Parse(other)

	if otherColor.IsEmpty() {
		return c
	}

	if c.IsEmpty() {
		return otherColor
	}

	// Calculate the new color
	return Color{
		Red:     mix(c.Red, otherColor.Red, percent),
		Green:   mix(c.Green, otherColor.Green, percent),
		Blue:    mix(c.Blue, otherColor.Blue, percent),
		Opacity: c.Opacity,
	}
}

// Text returns either black or white to use as text on top of the current color.
func (c Color) Text() Color {

	if c.IsEmpty() {
		return c
	}

	mean := c.Mean()

	if mean > 127 {
		return Black
	}

	return White
}

// TextExt uses the extended color tools to return a color that is suitable for text on top of the current color.
// It returns black or white for medium-range colors, and a blended value for very light and very dark colors.
func (c Color) TextExt() Color {

	if c.IsEmpty() {
		return c
	}

	mean := c.Mean()

	if mean > 212 {
		return c.Mix(Black, 80)
	}

	if mean > 169 {
		return c.Mix(Black, 90)
	}

	if mean > 127 {
		return Black
	}

	if mean > 85 {
		return White
	}

	if mean > 42 {
		return c.Mix(White, 90)
	}

	return c.Mix(White, 80)
}

// Contrast mixes this color with either black or white.  Light colors become darker,
// and dark colors become lighter.  The percentage determines how much to mix the colors.
func (c Color) Contrast(percent int) Color {

	if c.Mean() > 127 {
		return c.Darken(percent)
	}

	return c.Lighten(percent)
}

// Darken mixes this color with black, moving towards black by the specified percentage.
func (c Color) Darken(percent int) Color {
	return c.Mix(Black, percent)
}

// Lighten mixes this color with white, moving towards white by the specified percentage.
func (c Color) Lighten(percent int) Color {
	return c.Mix(White, percent)
}

func mix(color1 int, color2 int, percent int) int {

	if percent > 100 {
		percent = 100
	} else if percent < 0 {
		percent = 0
	}

	opacity2 := float64(percent) / 100
	opacity1 := 1 - opacity2

	return int((float64(color1)*opacity1 + float64(color2)*opacity2))
}
