package main

import "math"

type ray struct {
	origin    point3
	direction vec3
}

func (r *ray) at(t float64) point3 {
	d := Vec3_FMul(&r.direction, t)
	d.AddAssign(&r.origin)
	return *d
}

var bgColor1 color = color{1.0, 1.0, 1.0}
var bgColor2 color = color{0.5, 0.7, 1.0}

func Ray_Color(r *ray, world *hittable_list, depth int) *color {
	if depth < 0 {
		return &color{0, 0, 0}
	}
	hitSomething, rec := world.Hit(r, 0.001, math.Inf(1))
	if hitSomething {
		isScatter, attenuation, scattered := (*rec.matPtr).Scatter(r, rec)
		if isScatter {
			return Vec3_Mul(attenuation, Ray_Color(scattered, world, depth-1))
		}
	}
	t := 0.5 * (Vec3_UnitVector(&r.direction).y + 1.0)
	return Vec3_Add(
		Vec3_FMul(&bgColor1, (1.0-t)),
		Vec3_FMul(&bgColor2, t))
}
