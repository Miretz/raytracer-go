package main

import "math"

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
	*scattered = ray{rec.p, Vec3_AddMultiple(reflected, Vec3_FMul(&randomInUnit, l.fuzz))}
	*attenuation = l.albedo
	return Vec3_Dot(&scattered.direction, &rec.normal) > 0
}

type dielectric struct {
	ir float64
}

func (l *dielectric) Scatter(rayIn *ray, rec *hit_record, attenuation *color, scattered *ray) bool {
	*attenuation = color{1.0, 1.0, 1.0}
	refractionRatio := l.ir
	if rec.frontFace {
		refractionRatio = 1.0 / l.ir
	}
	unitDirection := Vec3_UnitVector(&rayIn.direction)
	unitDirNeg := unitDirection.Neg()
	cosTheta := math.Min(Vec3_Dot(&unitDirNeg, &rec.normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	cannotRefract := refractionRatio*sinTheta > 1.0

	direction := vec3{}
	if cannotRefract || l.reflectance(cosTheta, refractionRatio) > RandomFloat() {
		direction = Vec3_Reflect(&unitDirection, &rec.normal)
	} else {
		direction = Vec3_Refract(&unitDirection, &rec.normal, refractionRatio)
	}
	*scattered = ray{rec.p, direction}
	return true
}

func (l *dielectric) reflectance(cosine, refIdx float64) float64 {
	// Use Schlick's approximation for reflectance
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
