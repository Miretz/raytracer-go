package main

import "sort"

type hittable_list struct {
	objects []hittable
}

func (hl *hittable_list) Clear() {
	hl.objects = nil
}

func (hl *hittable_list) Add(object hittable) {
	hl.objects = append(hl.objects, object)
}

func (hl *hittable_list) GetCenter() point3 {
	return point3{0, 0, 0}
}

func (hl *hittable_list) SortByDistance(pos *vec3) {
	sort.Slice(hl.objects, func(p, q int) bool {
		aCenter := hl.objects[p].GetCenter()
		bCenter := hl.objects[q].GetCenter()
		aDist := Vec3_Sub(&aCenter, pos)
		bDist := Vec3_Sub(&bCenter, pos)
		return aDist.LengthSquared() < bDist.LengthSquared()
	})
}

func (hl *hittable_list) Hit(r *ray, tMin float64, tMax float64, rec *hit_record) bool {
	var tempRec hit_record
	hitAnything := false
	closestSoFar := tMax

	for i := 0; i < len(hl.objects); i++ {
		if hl.objects[i].Hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			*rec = tempRec
		}
	}

	return hitAnything
}
