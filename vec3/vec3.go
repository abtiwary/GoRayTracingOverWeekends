package vec3

import (
	"fmt"
	"math"
)

type Vec3 struct {
	Elements [3]float64
}

func NewVec3(e0, e1, e2 float64) *Vec3 {
	return &Vec3{
		Elements: [3]float64{e0, e1, e2},
	}
}

func (v *Vec3) X() float64 {
	return v.Elements[0]
}

func (v *Vec3) Y() float64 {
	return v.Elements[1]
}

func (v *Vec3) Z() float64 {
	return v.Elements[2]
}

func (v *Vec3) Copy(u *Vec3) {
	for i := 0; i < 3; i++ {
		v.Elements[i] = u.Elements[i]
	}
}

func (v *Vec3) Inverse() *Vec3 {
	return &Vec3{
		Elements: [3]float64{-v.X(), -v.Y(), -v.Z()},
	}
}

func (v *Vec3) AddVec3(v2 *Vec3) {
	v.Elements[0] += v2.Elements[0]
	v.Elements[1] += v2.Elements[1]
	v.Elements[2] += v2.Elements[2]
}

func (v *Vec3) MultBy(val float64) {
	v.Elements[0] *= val
	v.Elements[1] *= val
	v.Elements[2] *= val
}

func (v *Vec3) DivBy(val float64) {
	v.MultBy(1 / val)
}

func (v *Vec3) lengthSquared() float64 {
	return v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z()
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.lengthSquared())
}

func (v *Vec3) String() string {
	return fmt.Sprintf("%.4f %.4f %.4f", v.X(), v.Y(), v.Z())
}

// utility functions
func VectorAdd(u *Vec3, v *Vec3) *Vec3 {
	return NewVec3(
		u.X()+v.X(),
		u.Y()+v.Y(),
		u.Z()+v.Z(),
	)
}

func VectorSub(u *Vec3, v *Vec3) *Vec3 {
	return NewVec3(
		u.X()-v.X(),
		u.Y()-v.Y(),
		u.Z()-v.Z(),
	)
}

func VectorMult(u *Vec3, v *Vec3) *Vec3 {
	return NewVec3(
		u.X()*v.X(),
		u.Y()*v.Y(),
		u.Z()*v.Z(),
	)
}

func VectorMultBy(u *Vec3, val float64) *Vec3 {
	return NewVec3(
		u.X()*val,
		u.Y()*val,
		u.Z()*val,
	)
}

func VectorDivBy(u *Vec3, val float64) *Vec3 {
	return VectorMultBy(u, (1 / val))
}

func UnitVector(u *Vec3) *Vec3 {
	return VectorDivBy(u, u.Length())
}

func Dot(u *Vec3, v *Vec3) float64 {
	return u.X()*v.X() + u.Y()*v.Y() + u.Z()*v.Z()
}

func Cross(u *Vec3, v *Vec3) *Vec3 {
	return NewVec3(
		u.Y()*v.Z()-u.Z()*v.Y(),
		u.Z()*v.X()-u.X()*v.Z(),
		u.X()*v.Y()-u.Y()*v.X(),
	)
}
