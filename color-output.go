package color

import "strconv"

/******************************************
 * Output
 ******************************************/

// String returns this color as hex when it is fully opaque, as rgba() otherwise, and "" when it is empty
func (c Color) String() string {

	if c.IsEmpty() {
		return ""
	}

	if c.Opacity == 1 {
		return c.Hex()
	}

	return c.RGBA()
}

// Hex returns this color as a "#rrggbb" string
func (c Color) Hex() string {
	return "#" +
		formatHex(c.Red) +
		formatHex(c.Green) +
		formatHex(c.Blue)
}

// RGBA returns this color as an "rgba(r,g,b,a)" string
func (c Color) RGBA() string {
	return "rgba(" +
		strconv.Itoa(c.Red) + "," +
		strconv.Itoa(c.Green) + "," +
		strconv.Itoa(c.Blue) + "," +
		strconv.FormatFloat(c.Opacity, 'f', -1, 64) +
		")"
}
