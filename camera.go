package main

type camera struct {
	origin          point3
	lowerLeftCorner point3
	horizontal      vec3
	vertical        vec3
}

func NewCamera() camera {
	const aspectRatio = 16.0 / 9.0
	const viewportHeight = 2.0
	const viewportWidth = aspectRatio * viewportHeight
	const focalLength = 1.0

	c := camera{}
	c.origin = point3{0, 0, 0}
	c.horizontal = vec3{viewportWidth, 0, 0}
	c.vertical = vec3{0, viewportHeight, 0}
	c.lowerLeftCorner = c.origin.SubMultiple(
		c.horizontal.Div(2.0),
		c.vertical.Div(2.0),
		vec3{0, 0, focalLength})
	return c
}

func (c *camera) GetRay(u, v float64) ray {
	return ray{c.origin, Vec3_AddMultiple(
		c.lowerLeftCorner,
		c.horizontal.Mul(u),
		c.vertical.Mul(v),
		c.origin.Neg())}
}
