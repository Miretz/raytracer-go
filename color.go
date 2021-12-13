package main

import (
	"fmt"
	"math"
)

type color = vec3

func WriteColor(pixelColor *color, samplesPerPixel int) string {
	scale := 1.0 / float64(samplesPerPixel)
	r := math.Sqrt(scale * pixelColor.x)
	g := math.Sqrt(scale * pixelColor.y)
	b := math.Sqrt(scale * pixelColor.z)

	return fmt.Sprintf("%d %d %d\n",
		int32(256*Clamp(r, 0.0, 0.999)),
		int32(256*Clamp(g, 0.0, 0.999)),
		int32(256*Clamp(b, 0.0, 0.999)))
}
