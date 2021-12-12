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

func Ray_HitSphere(center *point3, radius float64, r *ray) float64 {
	oc := Vec3_Sub(&r.origin, center)
	a := r.direction.LengthSquared()
	halfB := Vec3_Dot(&oc, &r.direction)
	c := oc.LengthSquared() - radius*radius
	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return (-halfB - math.Sqrt(discriminant)) / a
	}
}

func Ray_Color(r *ray) color {
	t := Ray_HitSphere(&point3{0, 0, -1}, 0.5, r)
	if t > 0.0 {
		rayAtT := r.at(t)
		nInput := Vec3_Sub(&rayAtT, &vec3{0, 0, -1})
		N := Vec3_UnitVector(&nInput)
		colorBase := color{N.x + 1, N.y + 1, N.z + 1}
		return colorBase.Mul(0.5)
	}
	unitDirection := Vec3_UnitVector(&r.direction)
	t = 0.5 * (unitDirection.y + 1.0)
	c1 := Vec3_FMul(&color{1.0, 1.0, 1.0}, (1.0 - t))
	c2 := Vec3_FMul(&color{0.5, 0.7, 1.0}, t)
	return Vec3_Add(&c1, &c2)
}
