package color

type Color struct {
	Red     int
	Green   int
	Blue    int
	Opacity float64
}

func NewColor() Color {
	return Color{
		Red:     -1, // -1 means this color is empty
		Green:   -1, // -1 means this color is empty
		Blue:    -1, // -1 means this color is empty
		Opacity: 1,  // 1 means this color is fully opaque
	}
}

func Parse(value any) Color {

	switch typed := value.(type) {

	case Color:
		return typed

	case *Color:
		return *typed

	case string:

		if result := ParseHex(typed); result.NotEmpty() {
			return result
		}

		return ParseRGB(typed)
	}

	return NewColor()
}
