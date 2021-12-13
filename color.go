package main

import "fmt"

type color = vec3

func WriteColor(pixelColor *color, samplesPerPixel int) string {
	r := pixelColor.x
	g := pixelColor.y
	b := pixelColor.z

	scale := 1.0 / float64(samplesPerPixel)
	r *= scale
	g *= scale
	b *= scale

	return fmt.Sprintf("%d %d %d",
		int32(256*Clamp(r, 0.0, 0.999)),
		int32(256*Clamp(g, 0.0, 0.999)),
		int32(256*Clamp(b, 0.0, 0.999)))
}
