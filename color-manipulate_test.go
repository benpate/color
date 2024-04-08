package color

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMix_1(t *testing.T) {
	color := Parse("#000000").Mix("#ffffff", 50)
	require.Equal(t, "#7f7f7f", color.Hex())
}

func TestMix_2(t *testing.T) {
	color := Parse("#0000ff").Mix("#ffffff", 50)
	require.Equal(t, "#7f7fff", color.Hex())
}

func TestMix_3(t *testing.T) {
	color := Parse("#0000ff").Mix("#000000", 50)
	require.Equal(t, "#00007f", color.Hex())
}

func TestMix_4(t *testing.T) {
	color := Parse("#00ffff").Mix("#000000", 25)
	require.Equal(t, "#00bfbf", color.Hex())
}

func TestMix_5(t *testing.T) {
	color := Parse("#00ffff").Mix("#000000", 75)
	require.Equal(t, "#003f3f", color.Hex())
}

func TestLighten(t *testing.T) {
	color := Parse("#000000").Lighten(50)
	require.Equal(t, "#7f7f7f", color.Hex())
}

func TestLighten2(t *testing.T) {
	color := Parse("#000000").Lighten(12)
	require.Equal(t, "#1e1e1e", color.Hex())
}

func TestDarken(t *testing.T) {
	color := Parse("#ffffff").Darken(50)
	require.Equal(t, "#7f7f7f", color.Hex())
}

func TestDarken2(t *testing.T) {
	color := Parse("#ffffff").Darken(12)
	require.Equal(t, "#e0e0e0", color.Hex())
}
