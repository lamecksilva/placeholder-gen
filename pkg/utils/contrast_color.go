package utils

import (
	"image/color"
	"math"
)

func getRelativeLuminance(c color.RGBA) float64 {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0

	adjust := func(v float64) float64 {
		if v <= 0.03928 {
			return v / 12.92
		}
		return math.Pow((v+0.055)/1.055, 2.4)
	}

	return 0.2126*adjust(r) + 0.7152*adjust(g) + 0.0722*adjust(b)
}

func GetContrastColor(backgroundColor color.RGBA) color.RGBA {
	luminance := getRelativeLuminance(backgroundColor)

	if luminance > 0.5 {
		return color.RGBA{R: 0, G: 0, B: 0, A: 255}
	}
	return color.RGBA{R: 255, G: 255, B: 255, A: 255}
}

func ConvertToRGBA(c color.Color) color.RGBA {
	// Obtém os componentes RGBA no intervalo [0, 65535]
	r, g, b, a := c.RGBA()

	// Converte para o intervalo [0, 255] e retorna como color.RGBA
	return color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}
}
