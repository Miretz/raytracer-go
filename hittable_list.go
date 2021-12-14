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

func (hl *hittable_list) Hit(r *ray, tMin float64, tMax float64) (bool, *hit_record) {
	tempRec := new(hit_record)
	hitAnything := false
	closestSoFar := tMax

	for _, obj := range hl.objects {
		isHit, rec := obj.Hit(r, tMin, closestSoFar)
		if isHit {
			hitAnything = true
			closestSoFar = rec.t
			tempRec = rec
		}
	}

	return hitAnything, tempRec
}
