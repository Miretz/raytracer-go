package main

import (
	"math"
	"math/rand"
)

type material interface {
	Scatter(rayIn *ray, rec *hit_record, attenuation *color, scattered *ray) bool
}

type lambertian struct {
	albedo color
}

func (l *lambertian) Scatter(rayIn *ray, rec *hit_record, attenuation *color, scattered *ray) bool {
	scatterDirection := Vec3_AddMultiple(rec.normal, Vec3_RandomUnitVector())

	if scatterDirection.NearZero() {
		scatterDirection = rec.normal
	}

	*scattered = ray{rec.p, scatterDirection}
	*attenuation = l.albedo
	return true
}

type metal struct {
	albedo color
	fuzz   float64
}

func (l *metal) Scatter(rayIn *ray, rec *hit_record, attenuation *color, scattered *ray) bool {
	unitDir := Vec3_UnitVector(&rayIn.direction)
	reflected := Vec3_Reflect(&unitDir, &rec.normal)
	randomInUnit := Vec3_RandomInUnitSphere()
	randomInUnit.MulAssign(l.fuzz)
	*scattered = ray{rec.p, Vec3_AddMultiple(reflected, randomInUnit)}
	*attenuation = l.albedo
	return Vec3_Dot(&scattered.direction, &rec.normal) > 0
}

type dielectric struct {
	ir float64
}

var defaultGlassColor color = color{1.0, 1.0, 1.0}

func (l *dielectric) Scatter(rayIn *ray, rec *hit_record, attenuation *color, scattered *ray) bool {
	*attenuation = defaultGlassColor
	refractionRatio := l.ir
	if rec.frontFace {
		refractionRatio = 1.0 / l.ir
	}
	unitDirection := Vec3_UnitVector(&rayIn.direction)
	unitDirNeg := unitDirection.Neg()
	cosTheta := math.Min(Vec3_Dot(&unitDirNeg, &rec.normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	cannotRefract := refractionRatio*sinTheta > 1.0

	if cannotRefract || l.reflectance(cosTheta, refractionRatio) > rand.Float64() {
		*scattered = ray{rec.p, Vec3_Reflect(&unitDirection, &rec.normal)}
	} else {
		*scattered = ray{rec.p, Vec3_Refract(&unitDirection, &rec.normal, refractionRatio)}
	}
	return true
}

func (l *dielectric) reflectance(cosine, refIdx float64) float64 {
	// Use Schlick's approximation for reflectance
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
