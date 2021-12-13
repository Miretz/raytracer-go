package main

import "math"

type camera struct {
	origin          point3
	lowerLeftCorner point3
	horizontal      vec3
	vertical        vec3
}

func NewCamera(lookfrom point3, lookat point3, vup vec3,
	vFov float64, aspectRatio float64) camera {
	theta := DegreesToRadians(vFov)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	diffLook := Vec3_Sub(&lookfrom, &lookat)
	w := Vec3_UnitVector(&diffLook)
	cross := Vec3_Cross(&vup, &w)
	u := Vec3_UnitVector(&cross)
	v := Vec3_Cross(&w, &u)

	c := camera{}
	c.origin = lookfrom
	c.horizontal = Vec3_FMul(&u, viewportWidth)
	c.vertical = Vec3_FMul(&v, viewportHeight)
	c.lowerLeftCorner = c.origin.SubMultiple(
		c.horizontal.Div(2.0),
		c.vertical.Div(2.0),
		w)
	return c
}

func (c *camera) GetRay(s, t float64) ray {
	return ray{c.origin, Vec3_AddMultiple(
		c.lowerLeftCorner,
		c.horizontal.Mul(s),
		c.vertical.Mul(t),
		c.origin.Neg())}
}
