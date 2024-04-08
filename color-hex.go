package color

import (
	"strconv"
	"strings"
)

func ParseHex(value string) Color {

	result := NewColor()

	// Remove hash prefix (if present)
	value = strings.TrimPrefix(value, "#")

	// Require that the value is exactly six characters
	if len(value) != 6 {
		return result
	}

	// Split and parse
	result.Red = parseHex(value[0:2])
	result.Green = parseHex(value[2:4])
	result.Blue = parseHex(value[4:6])

	return result
}

// parseHex converts a hex string to an integer from 0-255.  If the
// string is not a valid hex value, or if the integer is out of range,
// -1 is returned.
func parseHex(value string) int {

	if result, err := strconv.ParseInt(value, 16, 64); err == nil {
		resultInt := int(result)
		if isValidInt(resultInt) {
			return resultInt
		}
	}

	return -1
}

// formatHex converts an integer from 0-255 to a two character hex string.  If the
// integer is outside that range, then "" is returned.
func formatHex(value int) string {

	if value < 0 {
		return ""
	}

	if value > 255 {
		return ""
	}

	result := strconv.FormatInt(int64(value), 16)

	if len(result) == 1 {
		result = "0" + result
	}
	return result
}
