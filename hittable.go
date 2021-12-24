package main

type hit_record struct {
	p         *point3
	normal    *vec3
	t         float64
	frontFace bool
	matPtr    *material
}

func (h *hit_record) SetFaceNormal(r *ray, outwardNormal *vec3) {
	h.frontFace = Vec3_Dot(&r.direction, outwardNormal) < 0
	if h.frontFace {
		h.normal = outwardNormal
	} else {
		h.normal = outwardNormal.Neg()
	}
}

type hittable interface {
	Hit(r *ray, tMin float64, tMax float64) (bool, *hit_record)
}
