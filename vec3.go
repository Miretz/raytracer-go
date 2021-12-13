package main

import (
	"fmt"
	"math"
)

type vec3 struct {
	x float64
	y float64
	z float64
}

func (a *vec3) Get(i int) float64 {
	switch i {
	case 0:
		return a.x
	case 1:
		return a.y
	case 2:
		return a.z
	default:
		return 0
	}
}

func (a *vec3) GetRef(i int) *float64 {
	switch i {
	case 0:
		return &a.x
	case 1:
		return &a.y
	case 2:
		return &a.z
	default:
		return nil
	}
}

func (a *vec3) Neg() vec3 {
	return vec3{-a.x, -a.y, -a.z}
}

func (a *vec3) Add(b *vec3) vec3 {
	return Vec3_Add(a, b)
}

func (a *vec3) Sub(b *vec3) vec3 {
	return Vec3_Sub(a, b)
}

func (a *vec3) AddMultiple(vecs ...vec3) vec3 {
	return Vec3_AddMultiple(append(vecs, *a)...)
}

func (a *vec3) SubMultiple(vecs ...vec3) vec3 {
	return Vec3_SubMultiple(a, vecs...)
}

func (a *vec3) Mul(t float64) vec3 {
	return vec3{a.x * t, a.y * t, a.z * t}
}

func (a *vec3) Div(t float64) vec3 {
	return a.Mul(1 / t)
}

func (a *vec3) Length() float64 {
	return math.Sqrt(a.LengthSquared())
}

func (a *vec3) LengthSquared() float64 {
	return (a.x * a.x) + (a.y * a.y) + (a.z * a.z)
}

func (a *vec3) NearZero() bool {
	const s = 1e-8
	return (math.Abs(a.x) < s) &&
		(math.Abs(a.y) < s) &&
		(math.Abs(a.z) < s)
}

// Utility functions

func Vec3_Print(v *vec3) string {
	return fmt.Sprintf("%4.f %4.f %4.f", v.x, v.y, v.z)
}

func Vec3_Add(a *vec3, b *vec3) vec3 {
	return vec3{
		a.x + b.x,
		a.y + b.y,
		a.z + b.z,
	}
}

func Vec3_AddMultiple(vecs ...vec3) vec3 {
	res := vec3{0, 0, 0}
	for _, v := range vecs {
		res = Vec3_Add(&res, &v)
	}
	return res
}

func Vec3_Sub(a *vec3, b *vec3) vec3 {
	return vec3{
		a.x - b.x,
		a.y - b.y,
		a.z - b.z,
	}
}

func Vec3_SubMultiple(original *vec3, vecs ...vec3) vec3 {
	res := *original
	for _, v := range vecs {
		res = Vec3_Sub(&res, &v)
	}
	return res
}

func Vec3_Mul(a *vec3, b *vec3) vec3 {
	return vec3{
		a.x * b.x,
		a.y * b.y,
		a.z * b.z,
	}
}

func Vec3_FMul(a *vec3, t float64) vec3 {
	return vec3{
		a.x * t,
		a.y * t,
		a.z * t,
	}
}

func Vec3_FDiv(a *vec3, t float64) vec3 {
	return Vec3_FMul(a, 1/t)
}

func Vec3_Dot(a *vec3, b *vec3) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func Vec3_Cross(u *vec3, v *vec3) vec3 {
	return vec3{
		u.Get(1)*v.Get(2) - u.Get(2)*v.Get(1),
		u.Get(2)*v.Get(0) - u.Get(0)*v.Get(2),
		u.Get(0)*v.Get(1) - u.Get(1)*v.Get(0),
	}
}

func Vec3_UnitVector(v *vec3) vec3 {
	return Vec3_FDiv(v, v.Length())
}

func Vec3_Random() vec3 {
	return vec3{RandomFloat(), RandomFloat(), RandomFloat()}
}

func Vec3_RandomBetween(min, max float64) vec3 {
	return vec3{
		RandomFloatBetween(min, max),
		RandomFloatBetween(min, max),
		RandomFloatBetween(min, max)}
}

func Vec3_RandomInUnitSphere() vec3 {
	for {
		p := Vec3_RandomBetween(-1.0, 1.0)
		if p.LengthSquared() < 1.0 {
			return p
		}
	}
}

func Vec3_RandomUnitVector() vec3 {
	t := Vec3_RandomInUnitSphere()
	return Vec3_UnitVector(&t)
}

func Vec3_RandomInHemisphere(normal *vec3) vec3 {
	inUnitSphere := Vec3_RandomInUnitSphere()
	if Vec3_Dot(&inUnitSphere, normal) > 0.0 {
		return inUnitSphere
	}
	return inUnitSphere.Neg()
}

func Vec3_Reflect(v, n *vec3) vec3 {
	return Vec3_SubMultiple(v, Vec3_FMul(n, 2.0*Vec3_Dot(v, n)))
}

func Vec3_Refract(uv *vec3, n *vec3, etaiOverEtat float64) vec3 {
	negUv := uv.Neg()
	cosTheta := math.Min(Vec3_Dot(&negUv, n), 1.0)
	tmp := Vec3_AddMultiple(*uv, Vec3_FMul(n, cosTheta))
	rOutPerp := Vec3_FMul(&tmp, etaiOverEtat)
	parallelMul := -math.Sqrt(math.Abs(1.0 - rOutPerp.LengthSquared()))
	rOutParallel := Vec3_FMul(n, parallelMul)
	return Vec3_Add(&rOutPerp, &rOutParallel)
}

type point3 = vec3
