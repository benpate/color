package color

import (
	"strconv"
	"strings"
)

// parseFloat converts a number string to a float from 0-1.  If the
// string is not a valid float, or if the float is out of range,
// -1 is returned.
func parseFloat(value string) float64 {

	// RULE: CSS permits spaces around the separators (`rgba(0, 0, 0, 0.5)`),
	// and ParseFloat rejects them, so the value must be trimmed before parsing.
	value = strings.TrimSpace(value)

	if result, err := strconv.ParseFloat(value, 64); err == nil {
		if result >= 0 && result <= 1 {
			return result
		}
	}

	return -1
}

// isValidInt returns true if the value is between 0 and 255, inclusive.
func isValidInt(value int) bool {
	return value >= 0 && value <= 255
}
