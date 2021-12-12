package main

import "fmt"

type color = vec3

func WriteColor(color *color) string {
	return fmt.Sprintf("%d %d %d",
		int32(color.x*255.999),
		int32(color.y*255.999),
		int32(color.z*255.999))
}
