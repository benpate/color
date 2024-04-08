package color

import "strconv"

/******************************************
 * Output
 ******************************************/

func (c Color) String() string {

	if c.IsEmpty() {
		return ""
	}

	if c.Opacity == 1 {
		return c.Hex()
	}

	return c.RGBA()
}

func (c Color) Hex() string {
	return "#" +
		formatHex(c.Red) +
		formatHex(c.Green) +
		formatHex(c.Blue)
}

func (c Color) RGBA() string {
	return "rgba(" +
		strconv.Itoa(c.Red) + "," +
		strconv.Itoa(c.Green) + "," +
		strconv.Itoa(c.Blue) + "," +
		strconv.FormatFloat(c.Opacity, 'f', -1, 64) +
		")"
}
