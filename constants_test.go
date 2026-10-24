package color

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestConstants verifies that the named colors match their hex strings
func TestConstants(t *testing.T) {

	require.Equal(t, White, ParseHex(HexWhite))
	require.Equal(t, Black, ParseHex(HexBlack))
	require.Equal(t, "#ffffff", White.String())
	require.Equal(t, "#000000", Black.String())
}
