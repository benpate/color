// Package color parses, inspects, mixes, and formats the RGB colors used in web themes.
//
// Parse reads a Color from a "#rrggbb" hex string, an "rgb(r, g, b)" or "rgba(r, g, b, a)"
// string, a Color, or a *Color. A value it cannot read becomes an empty Color: at least one
// channel is -1, IsEmpty reports TRUE, and String formats it as "".
//
// A Color can report whether it is light or dark, mix toward another color by a percentage,
// and choose a text color that stays readable on top of it. String writes a fully opaque
// color as hex and any other color as rgba().
package color
