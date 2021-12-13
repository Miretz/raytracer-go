package main

import "math/rand"

const Pi = 3.1415926535897932385

func DegreesToRadians(degrees float64) float64 {
	return degrees * Pi / 180.0
}

func RandomFloat() float64 {
	return rand.Float64()
}

func RandomFloatBetween(min, max float64) float64 {
	return min + (max-min)*RandomFloat()
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
