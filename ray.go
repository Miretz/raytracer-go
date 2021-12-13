package main

import "math"

type ray struct {
	origin    point3
	direction vec3
}

func (r *ray) at(t float64) point3 {
	d := Vec3_FMul(&r.direction, t)
	d.AddAssign(&r.origin)
	return d
}

var bgColor1 color = color{1.0, 1.0, 1.0}
var bgColor2 color = color{0.5, 0.7, 1.0}
var emptyColor color = color{0, 0, 0}

func Ray_Color(r *ray, world *hittable_list, depth int) color {
	if depth < 0 {
		return emptyColor
	}

	var rec hit_record
	if world.Hit(r, 0.001, math.Inf(1), &rec) {
		scattered := ray{}
		attenuation := color{}
		if (*rec.matPtr).Scatter(r, &rec, &attenuation, &scattered) {
			rc := Ray_Color(&scattered, world, depth-1)
			return Vec3_Mul(&attenuation, &rc)
		}
	}
	unitDirection := Vec3_UnitVector(&r.direction)
	t := 0.5 * (unitDirection.y + 1.0)
	r1 := Vec3_FMul(&bgColor1, (1.0 - t))
	r2 := Vec3_FMul(&bgColor2, t)
	r1.AddAssign(&r2)
	return r1
}
