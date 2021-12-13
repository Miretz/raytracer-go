package main

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
