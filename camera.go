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

	diffLook := Vec3_Sub(&lookfrom, &lookat)
	c.w = Vec3_UnitVector(&diffLook)
	cross := Vec3_Cross(&vup, &c.w)
	c.u = Vec3_UnitVector(&cross)
	c.v = Vec3_Cross(&c.w, &c.u)

	c.origin = lookfrom
	c.horizontal = c.u.Mul(viewportWidth * focusDist)
	c.vertical = c.v.Mul(viewportHeight * focusDist)
	c.lowerLeftCorner = c.origin.SubMultiple(
		c.horizontal.Div(2.0),
		c.vertical.Div(2.0),
		c.w.Mul(focusDist))
	c.lensRadius = aperture / 2
	return c
}

func (c *camera) GetRay(s, t float64) ray {
	rndDisk := Vec3_RandomInUnitDisk()
	rd := rndDisk.Mul(c.lensRadius)
	offset := Vec3_AddMultiple(
		c.u.Mul(rd.x),
		c.v.Mul(rd.y))
	return ray{
		c.origin.Add(&offset),
		Vec3_AddMultiple(
			c.lowerLeftCorner,
			c.horizontal.Mul(s),
			c.vertical.Mul(t),
			c.origin.Neg(),
			offset.Neg())}

}
