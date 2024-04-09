package color

/******************************************
 * Color Calculations
 ******************************************/

// IsEmpty returns TRUE if this color is empty.  This represents a Nil color
func (c Color) IsEmpty() bool {

	if c.Red == -1 {
		return true
	}

	if c.Green == -1 {
		return true
	}

	if c.Blue == -1 {
		return true
	}

	return false
}

// NotEmpty returns TRUE if this color is NOT empty
func (color Color) NotEmpty() bool {
	return !color.IsEmpty()
}

// Mean returns the average of the red, green, and blue values.
func (c Color) Mean() int {

	if c.IsEmpty() {
		return -1
	}

	return (c.Red + c.Green + c.Blue) / 3
}

// ColorMode returns "light" if this is a light color, or "dark" if this is a dark color.
func (c Color) ColorMode() string {

	if c.IsEmpty() {
		return ""
	}

	if c.Mean() > 127 {
		return "light"
	}

	return "dark"
}

func (c Color) IsDark() bool {
	return c.ColorMode() == "dark"
}

func (c Color) IsLight() bool {
	return c.ColorMode() == "light"
}
