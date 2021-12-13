package main

type hittable_list struct {
	objects []hittable
}

func (hl *hittable_list) Clear() {
	hl.objects = nil
}

func (hl *hittable_list) Add(object hittable) {
	hl.objects = append(hl.objects, object)
}

func (hl *hittable_list) Hit(r *ray, tMin float64, tMax float64, rec *hit_record) bool {
	tempRec := hit_record{}
	hitAnything := false
	closestSoFar := tMax

	for _, object := range hl.objects {
		if object.Hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			*rec = tempRec
		}
	}

	return hitAnything
}
