package color

// HexWhite is white, as a hex string
const HexWhite = "#FFFFFF"

// HexBlack is black, as a hex string
const HexBlack = "#000000"

// ColorModeLight names a light color mode
const ColorModeLight = "light"

// ColorModeDark names a dark color mode
const ColorModeDark = "dark"

// ColorModeAuto names a color mode that follows the viewer's preference
const ColorModeAuto = "auto"

// ColorModeExtLight names the lightest of the four extended color modes
const ColorModeExtLight = "light"

// ColorModeExtMediumLight names the second-lightest extended color mode
const ColorModeExtMediumLight = "medium-light"

// ColorModeExtMediumDark names the second-darkest extended color mode
const ColorModeExtMediumDark = "medium-dark"

// ColorModeExtDark names the darkest of the four extended color modes
const ColorModeExtDark = "dark"

// White is fully opaque white
var White = Color{
	Red:     255,
	Green:   255,
	Blue:    255,
	Opacity: 1.0,
}

// Black is fully opaque black
var Black = Color{
	Red:     0,
	Green:   0,
	Blue:    0,
	Opacity: 1.0,
}
