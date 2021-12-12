package main

type ray struct {
	origin    point3
	direction vec3
}

func (r *ray) at(t float64) point3 {
	d := Vec3_FMul(&r.direction, t)
	return Vec3_Add(&r.origin, &d)
}

func Ray_Color(r *ray) color {
	unitDirection := Vec3_UnitVector(&r.direction)
	t := 0.5 * (unitDirection.y + 1.0)
	c1 := Vec3_FMul(&color{1.0, 1.0, 1.0}, (1.0 - t))
	c2 := Vec3_FMul(&color{0.5, 0.7, 1.0}, t)
	return Vec3_Add(&c1, &c2)
}
