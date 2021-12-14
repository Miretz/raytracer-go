package main

import "math"

type camera struct {
	origin          point3
	lowerLeftCorner point3
	horizontal      vec3
	vertical        vec3
	u               vec3
	v               vec3
	w               vec3
	lensRadius      float64
}

func NewCamera(lookfrom point3, lookat point3, vup vec3,
	vFov float64, aspectRatio float64,
	aperture float64, focusDist float64) camera {
	theta := DegreesToRadians(vFov)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	c := camera{}

	c.w = *Vec3_UnitVector(Vec3_Sub(&lookfrom, &lookat))
	c.u = *Vec3_UnitVector(Vec3_Cross(&vup, &c.w))
	c.v = *Vec3_Cross(&c.w, &c.u)

	c.origin = lookfrom
	c.horizontal = *Vec3_FMul(&c.u, viewportWidth*focusDist)
	c.vertical = *Vec3_FMul(&c.v, viewportHeight*focusDist)
	c.lowerLeftCorner = *Vec3_SubMultiple(
		&c.origin,
		Vec3_FDiv(&c.horizontal, 2.0),
		Vec3_FDiv(&c.vertical, 2.0),
		Vec3_FMul(&c.w, focusDist))
	c.lensRadius = aperture / 2
	return c
}

func (c *camera) GetRay(s, t float64) *ray {
	rd := Vec3_FMul(Vec3_RandomInUnitDisk(), c.lensRadius)
	offset := Vec3_Add(
		Vec3_FMul(&c.u, rd.x),
		Vec3_FMul(&c.v, rd.y))
	return &ray{
		*Vec3_Add(&c.origin, offset),
		*Vec3_AddMultiple(
			&c.lowerLeftCorner,
			Vec3_FMul(&c.horizontal, s),
			Vec3_FMul(&c.vertical, t),
			c.origin.Neg(),
			offset.Neg())}
}
