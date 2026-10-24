package color

// Color is an RGB color with an opacity. A channel of -1 marks the color as empty.
type Color struct {
	Red     int
	Green   int
	Blue    int
	Opacity float64
}

// Parse converts a Color, a *Color, or a hex or rgb()/rgba() string into a Color.
// Anything else, including a nil *Color, returns an empty Color.
func Parse(value any) Color {

	switch typed := value.(type) {

	case Color:
		return typed

	case *Color:

		// RULE: A nil pointer is an empty color, not a panic
		if typed == nil {
			return NewColor()
		}

		return *typed

	case string:

		if result := ParseHex(typed); result.NotEmpty() {
			return result
		}

		return ParseRGB(typed)
	}

	return NewColor()
}

// NewColor returns an empty, fully opaque Color
func NewColor() Color {
	return Color{
		Red:     -1, // -1 means this color is empty
		Green:   -1, // -1 means this color is empty
		Blue:    -1, // -1 means this color is empty
		Opacity: 1,  // 1 means this color is fully opaque
	}
}
