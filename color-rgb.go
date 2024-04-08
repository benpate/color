package color

import (
	"strconv"
	"strings"
)

func ParseRGB(value string) Color {

	result := NewColor()
	hasOpacity := false

	if strings.HasPrefix(value, "rgb(") {
		value = strings.TrimPrefix(value, "rgb(")
		hasOpacity = false
	} else if strings.HasPrefix(value, "rgba(") {
		value = strings.TrimPrefix(value, "rgba(")
		hasOpacity = true
	} else {
		return result
	}

	value = strings.TrimSuffix(value, ")")

	parts := strings.Split(value, ",")
	partsLen := len(parts)

	if hasOpacity {
		if partsLen == 4 {
			result.Opacity = parseFloat(parts[3])
		} else {
			return result
		}
	} else {
		if partsLen != 3 {
			return result
		}
	}

	result.Red = parseInt(parts[0])
	result.Green = parseInt(parts[1])
	result.Blue = parseInt(parts[2])

	return result
}

// parseInt converts a number string to an integer from 0-255.  If the
// string is not a valid integer, or if the integer is out of range,
// -1 is returned.
func parseInt(value string) int {

	if result, err := strconv.ParseInt(value, 10, 8); err == nil {
		resultInt := int(result)
		if isValidInt(resultInt) {
			return resultInt
		}
	}

	return -1
}
