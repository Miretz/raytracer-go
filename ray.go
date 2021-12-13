package main

import "math"

type ray struct {
	origin    point3
	direction vec3
}

func (r *ray) at(t float64) point3 {
	d := r.direction.Mul(t)
	return Vec3_Add(&r.origin, &d)
}

func Ray_Color(r *ray, world *hittable_list, depth int) color {
	rec := hit_record{}

	if depth < 0 {
		return color{0, 0, 0}
	}

	if (*world).Hit(r, 0.001, math.Inf(1), &rec) {
		scattered := ray{}
		attenuation := color{}
		if (*rec.matPtr).Scatter(r, &rec, &attenuation, &scattered) {
			rc := Ray_Color(&scattered, world, depth-1)
			return Vec3_Mul(&attenuation, &rc)
		}
	}
	unitDirection := Vec3_UnitVector(&r.direction)
	t := 0.5 * (unitDirection.y + 1.0)
	r1 := Vec3_FMul(&color{1.0, 1.0, 1.0}, (1.0 - t))
	r2 := Vec3_FMul(&color{0.5, 0.7, 1.0}, t)
	return Vec3_Add(&r1, &r2)
}
