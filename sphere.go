package main

import "math"

type sphere struct {
	center point3
	radius float64
	matPtr *material
}

func (s *sphere) Hit(r *ray, tMin float64, tMax float64, rec *hit_record) bool {
	oc := Vec3_Sub(&r.origin, &s.center)
	a := r.direction.LengthSquared()
	halfB := Vec3_Dot(&oc, &r.direction)
	c := oc.LengthSquared() - s.radius*s.radius
	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)

	// find the nearest root that lies in the acceptable range
	root := (-halfB - sqrtd) / a
	if root < tMin || tMax < root {
		root = (-halfB + sqrtd) / a
		if root < tMin || tMax < root {
			return false
		}
	}
	rec.t = root
	rec.p = r.at(rec.t)
	temp := Vec3_Sub(&rec.p, &s.center)
	outwardNormal := temp.Div(s.radius)
	rec.SetFaceNormal(r, &outwardNormal)
	rec.matPtr = s.matPtr
	return true
}
