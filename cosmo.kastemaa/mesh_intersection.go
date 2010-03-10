package main

import "fmt"
import "./vec"

func main() {
	mesh := Cube(1) // a unit cube
	ray := Ray{vec.Make(1, 1, 2), vec.Make(0, 0, -1)}
	fmt.Printf("%v\n", ray.IntersectsMesh(mesh))
}

func Cube(radius float64) *Mesh {
	sides := make([]*Mesh, 6)
	i := 0
	for _, r := range []float64{radius, -radius} {
		// top
		sides[i] = Quadrilateral(
			vec.Make(-r, r, -r),
			vec.Make(r, r, -r),
			vec.Make(r, r, r),
			vec.Make(-r, r, r))
		i++
		// front
		sides[i] = Quadrilateral(
			vec.Make(-r, r, r),
			vec.Make(r, r, r),
			vec.Make(r, -r, r),
			vec.Make(-r, -r, r))
		i++
		// right
		sides[i] = Quadrilateral(
			vec.Make(r, r, r),
			vec.Make(r, r, -r),
			vec.Make(r, -r, -r),
			vec.Make(r, -r, r))
		i++
	}
	return Fuse(sides)
}

func Quadrilateral(nw *vec.Vector, ne *vec.Vector, se *vec.Vector, sw *vec.Vector) *Mesh {
	tri1 := Triangle{nw, ne, se}
	tri2 := Triangle{se, sw, nw}
	return &Mesh{[]*Triangle{&tri1, &tri2}}
}

func Fuse(meshes []*Mesh) *Mesh {
	size := 0
	for _, mesh := range meshes {
		size += len(mesh.Triangles)
	}
	triangles := make([]*Triangle, size)
	i := 0
	for _, mesh := range meshes {
		for _, tri := range mesh.Triangles {
			triangles[i] = tri
			i++
		}
	}
	return &Mesh{triangles}
}

type Ray struct {
	Start *vec.Vector
	Direction *vec.Vector // relative direction
}

func (ray *Ray) String() string {
	return fmt.Sprintf("Ray %v %v", ray.Start, ray.Direction)
}

func (ray *Ray) IntersectsTriangle(tri *Triangle) bool {
	u := vec.Subtract(tri.B, tri.A)
	v := vec.Subtract(tri.C, tri.A)
	if vec.Parallel(u, v) { // the triangle is degenerate
		// we could check for a ray/line-segment intersection here, but it seems pointless
		return false
	}
	// n is the normal to the supporting plane of the triangle
	n := vec.CrossProduct(u, v)
	// w0 is a relative vector from the start of the ray to one corner of the triangle
	w0 := vec.Subtract(ray.Start, tri.A)	
	if vec.Orthogonal(n, ray.Direction) { // the ray is parallel to the triangle
		// again, ray/line-segment intersection
		return false
	} else {
		a := -vec.DotProduct(n, w0)
		b := vec.DotProduct(n, ray.Direction)
		r := a / b
		if r < 0 {
			return false
		}
		i := vec.Add(ray.Start, vec.Scale(r, ray.Direction))
		uu := vec.DotProduct(u, u)
		uv := vec.DotProduct(u, v)
		vv := vec.DotProduct(v, v)
		w := vec.Subtract(i, tri.A)
		wu := vec.DotProduct(w, u)
		wv := vec.DotProduct(w, v)
		d := uv * uv - uu * vv
		s := (uv * wv - vv * wu) / d
		if s < 0 || s > 1 {
			return false
		}
		t := (uv * wu - uu * wv) / d
		if t < 0 || s + t > 1 {
			return false
		}
	}
	return true
}

func (ray *Ray) IntersectsMesh(mesh *Mesh) bool {
	for _, tri := range mesh.Triangles {
		if ray.IntersectsTriangle(tri) {
			return true
		}
	}
	return false
}

type Mesh struct {
	Triangles []*Triangle
}

func (mesh *Mesh) String() string {
	listStr := ""
	for _, tri := range mesh.Triangles {
		if len(listStr) > 0 {
			listStr += ", "
		}
		listStr += tri.String()
	}
	return fmt.Sprintf("Mesh {%v}", listStr)
}

type Triangle struct {
	A *vec.Vector
	B *vec.Vector
	C *vec.Vector
}

func (tri *Triangle) String() string {
	return fmt.Sprintf("Triangle %v %v %v", tri.A, tri.B, tri.C)
}
