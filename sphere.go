package main

import "math"

type sphere struct {
	center point3
	radius float64
	matPtr *material
}

func (s *sphere) GetCenter() point3 {
	return s.center
}

func (s *sphere) Hit(r *ray, tMin float64, tMax float64, rec *hit_record) bool {
	oc := Vec3_Sub(&r.origin, &s.center)
	halfB := Vec3_Dot(&oc, &r.direction)
	c := Vec3_LengthSquared(&oc) - s.radius*s.radius
	a := Vec3_LengthSquared(&r.direction)
	discriminant := halfB*halfB - a*c
	if discriminant < 0.0 {
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
	outwardNormal := Vec3_Sub(&rec.p, &s.center)
	outwardNormal.DivAssign(s.radius)
	rec.SetFaceNormal(r, &outwardNormal)
	rec.matPtr = s.matPtr
	return true
}
