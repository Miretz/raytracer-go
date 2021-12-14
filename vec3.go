package main

import (
	"math"
	"math/rand"
)

type vec3 struct {
	x float64
	y float64
	z float64
}

func (a *vec3) Neg() *vec3 {
	return &vec3{-a.x, -a.y, -a.z}
}

func (a *vec3) AddAssign(b *vec3) {
	a.x += b.x
	a.y += b.y
	a.z += b.z
}

func (a *vec3) SubAssign(b *vec3) {
	a.x -= b.x
	a.y -= b.y
	a.z -= b.z
}

func (a *vec3) MulAssign(t float64) {
	a.x *= t
	a.y *= t
	a.z *= t
}

func (a *vec3) DivAssign(t float64) {
	a.MulAssign(1.0 / t)
}

func (a *vec3) Length() float64 {
	return math.Sqrt(a.LengthSquared())
}

func (a *vec3) LengthSquared() float64 {
	return a.x*a.x + a.y*a.y + a.z*a.z
}

func (a *vec3) NearZero() bool {
	const s = 1e-8
	return (math.Abs(a.x) < s) &&
		(math.Abs(a.y) < s) &&
		(math.Abs(a.z) < s)
}

// Utility functions

func Vec3_Sub(a *vec3, b *vec3) *vec3 {
	return &vec3{
		a.x - b.x,
		a.y - b.y,
		a.z - b.z}
}

func Vec3_Add(a *vec3, b *vec3) *vec3 {
	return &vec3{
		a.x + b.x,
		a.y + b.y,
		a.z + b.z}
}

func Vec3_AddMultiple(vecs ...*vec3) *vec3 {
	res := vec3{0, 0, 0}
	for _, v := range vecs {
		res.AddAssign(v)
	}
	return &res
}

func Vec3_SubMultiple(original *vec3, vecs ...*vec3) *vec3 {
	res := *original
	for _, v := range vecs {
		res.SubAssign(v)
	}
	return &res
}

func Vec3_Mul(a *vec3, b *vec3) *vec3 {
	return &vec3{
		a.x * b.x,
		a.y * b.y,
		a.z * b.z,
	}
}

func Vec3_FMul(a *vec3, t float64) *vec3 {
	return &vec3{
		a.x * t,
		a.y * t,
		a.z * t,
	}
}

func Vec3_FDiv(a *vec3, t float64) *vec3 {
	return Vec3_FMul(a, 1/t)
}

func Vec3_Dot(a *vec3, b *vec3) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func Vec3_Cross(u *vec3, v *vec3) *vec3 {
	return &vec3{
		u.y*v.z - u.z*v.y,
		u.z*v.x - u.x*v.z,
		u.x*v.y - u.y*v.x,
	}
}

func Vec3_LengthSquared(a *vec3) float64 {
	return a.x*a.x + a.y*a.y + a.z*a.z
}

func Vec3_UnitVector(v *vec3) *vec3 {
	return Vec3_FDiv(v, v.Length())
}

func Vec3_Random() *vec3 {
	return &vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func Vec3_RandomBetween(min, max float64) *vec3 {
	return &vec3{
		RandomFloatBetween(min, max),
		RandomFloatBetween(min, max),
		RandomFloatBetween(min, max)}
}

func Vec3_RandomInUnitSphere() *vec3 {
	for {
		p := Vec3_RandomBetween(-1.0, 1.0)
		if p.LengthSquared() < 1.0 {
			return p
		}
	}
}

func Vec3_RandomUnitVector() *vec3 {
	return Vec3_UnitVector(Vec3_RandomInUnitSphere())
}

func Vec3_RandomInHemisphere(normal *vec3) *vec3 {
	inUnitSphere := Vec3_RandomInUnitSphere()
	if Vec3_Dot(inUnitSphere, normal) > 0.0 {
		return inUnitSphere
	}
	return inUnitSphere.Neg()
}

func Vec3_Reflect(v, n *vec3) *vec3 {
	t := Vec3_FMul(n, 2.0*Vec3_Dot(v, n))
	return Vec3_Sub(v, t)
}

func Vec3_Refract(uv *vec3, n *vec3, etaiOverEtat float64) *vec3 {
	cosTheta := math.Min(Vec3_Dot(uv.Neg(), n), 1.0)
	t := Vec3_FMul(n, cosTheta)
	rOutPerp := Vec3_FMul(Vec3_Add(uv, t), etaiOverEtat)
	parallelMul := -math.Sqrt(math.Abs(1.0 - rOutPerp.LengthSquared()))
	rOutParallel := Vec3_FMul(n, parallelMul)
	rOutPerp.AddAssign(rOutParallel)
	return rOutPerp
}

func Vec3_RandomInUnitDisk() *vec3 {
	for {
		p := vec3{
			RandomFloatBetween(-1.0, 1.0),
			RandomFloatBetween(-1.0, 1.0),
			0.0}
		if p.LengthSquared() < 1.0 {
			return &p
		}
	}
}

type point3 = vec3
