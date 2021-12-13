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

func Ray_Color(r *ray, world *hittable_list) color {
	rec := hit_record{}
	if (*world).Hit(r, 0, math.Inf(1), &rec) {
		temp := Vec3_Add(&rec.normal, &color{1, 1, 1})
		return Vec3_FMul(&temp, 0.5)
	}
	unitDirection := Vec3_UnitVector(&r.direction)
	t := 0.5 * (unitDirection.y + 1.0)
	r1 := Vec3_FMul(&color{1.0, 1.0, 1.0}, (1.0 - t))
	r2 := Vec3_FMul(&color{0.5, 0.7, 1.0}, t)
	return Vec3_Add(&r1, &r2)
}
