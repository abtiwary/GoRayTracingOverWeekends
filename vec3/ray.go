package vec3

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func NewRay(point *Vec3, dir *Vec3) *Ray {
	ray := &Ray{}
	if point != nil {
		ray.Origin.Copy(point)
	}
	if dir != nil {
		ray.Direction.Copy(dir)
	}
	return ray
}

func (r *Ray) At(t float64) *Vec3 {
	return VectorAdd(&r.Origin, VectorMultBy(&r.Direction, t))
}
