package color

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestString verifies hex for opaque colors, rgba() for translucent ones, and "" for empty ones
func TestString(t *testing.T) {

	require.Equal(t, "#010203", Color{Red: 1, Green: 2, Blue: 3, Opacity: 1}.String())
	require.Equal(t, "rgba(1,2,3,0.5)", Color{Red: 1, Green: 2, Blue: 3, Opacity: 0.5}.String())
	require.Equal(t, "rgba(1,2,3,0)", Color{Red: 1, Green: 2, Blue: 3, Opacity: 0}.String())
	require.Equal(t, "", NewColor().String())
}

// TestHex verifies two lowercase hex digits per channel
func TestHex(t *testing.T) {

	require.Equal(t, "#000000", Black.Hex())
	require.Equal(t, "#ffffff", White.Hex())
	require.Equal(t, "#010aff", Color{Red: 1, Green: 10, Blue: 255, Opacity: 1}.Hex())

	// Hex does not check for an empty or out-of-range color, so each bad channel writes nothing
	require.Equal(t, "#", NewColor().Hex())
	require.Equal(t, "#0000", Color{Red: 300, Green: 0, Blue: 0, Opacity: 1}.Hex())
}

// TestRGBA verifies the rgba() format, which RGBA writes even for an empty color
func TestRGBA(t *testing.T) {

	require.Equal(t, "rgba(1,2,3,0.5)", Color{Red: 1, Green: 2, Blue: 3, Opacity: 0.5}.RGBA())
	require.Equal(t, "rgba(1,2,3,1)", Color{Red: 1, Green: 2, Blue: 3, Opacity: 1}.RGBA())
	require.Equal(t, "rgba(1,2,3,0.125)", Color{Red: 1, Green: 2, Blue: 3, Opacity: 0.125}.RGBA())
	require.Equal(t, "rgba(-1,-1,-1,1)", NewColor().RGBA())
}
