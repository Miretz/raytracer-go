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

func (a *vec3) Neg() vec3 {
	a.x = -a.x
	a.y = -a.y
	a.z = -a.z
	return *a
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

func (a *vec3) Add(b *vec3) vec3 {
	a.x += b.x
	a.y += b.y
	a.z += b.z
	return *a
}

func (a *vec3) Sub(b *vec3) vec3 {
	a.x -= b.x
	a.y -= b.y
	a.z -= b.z
	return *a
}

func (a *vec3) Mul(t float64) vec3 {
	a.x *= t
	a.y *= t
	a.z *= t
	return *a
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

func Vec3_Sub(a *vec3, b *vec3) vec3 {
	return vec3{
		a.x - b.x,
		a.y - b.y,
		a.z - b.z,
	}
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
		u.Get(1)*v.Get(1) - u.Get(2)*v.Get(1),
		u.Get(2)*v.Get(0) - u.Get(0)*v.Get(2),
		u.Get(0)*v.Get(1) - u.Get(1)*v.Get(0),
	}
}

func Vec3_UnitVector(v *vec3) vec3 {
	return Vec3_FDiv(v, v.Length())
}
