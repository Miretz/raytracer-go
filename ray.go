package main

import "math"

type ray struct {
	origin    point3
	direction vec3
}

func (r *ray) at(t float64) point3 {
	d := Vec3_FMul(&r.direction, t)
	return Vec3_Add(&r.origin, &d)
}

func Ray_Color(r *ray, world *hittable_list, depth int) color {
	rec := hit_record{}

	if depth < 0 {
		return color{0, 0, 0}
	}

	if (*world).Hit(r, 0.001, math.Inf(1), &rec) {
		//target := Vec3_AddMultiple(rec.p, rec.normal, Vec3_RandomUnitVector())
		target := Vec3_AddMultiple(rec.p, Vec3_RandomInHemisphere(&rec.normal))
		newRay := ray{rec.p, Vec3_Sub(&target, &rec.p)}
		rc := Ray_Color(&newRay, world, depth-1)
		return Vec3_FMul(&rc, 0.5)
	}
	unitDirection := Vec3_UnitVector(&r.direction)
	t := 0.5 * (unitDirection.y + 1.0)
	r1 := Vec3_FMul(&color{1.0, 1.0, 1.0}, (1.0 - t))
	r2 := Vec3_FMul(&color{0.5, 0.7, 1.0}, t)
	return Vec3_Add(&r1, &r2)
}
