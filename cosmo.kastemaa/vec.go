package vec

import "fmt"
import "math"

type Vector struct {
	X float64
	Y float64
	Z float64
}

func Make(x float64, y float64, z float64) *Vector {
	return &Vector{x, y, z}
}

func (v *Vector) String() string {
	return fmt.Sprintf("[%v %v %v]", v.X, v.Y, v.Z)
}

func AlmostEqual(v *Vector, w *Vector) bool {
	return fAlmostEqual(v.X, w.X) && fAlmostEqual(v.Y, w.Y) && fAlmostEqual(v.Z, w.Z)
}

func Magnitude(v *Vector) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func Unit(v *Vector) *Vector {
	mag := Magnitude(v)
	return Make(v.X / mag, v.Y / mag, v.Z / mag)
}

func Scale(s float64, v *Vector) *Vector {
	return Make(s * v.X, s * v.X, s * v.X)
}

func CrossProduct(v *Vector, w *Vector) *Vector {
    x := v.Y * w.Z - v.Z * w.Y
    y := v.Z * w.X - v.X * w.Z
    z := v.X * w.Y - v.Y * w.X
    return Make(x, y, z)
}

func Add(v *Vector, w *Vector) *Vector {
	return Make(v.X + w.X, v.Y + w.Y, v.Z + w.Z)
}

func Negate(v *Vector) *Vector {
	return Make(-v.X, -v.Y, -v.Z)
}

func Subtract(v *Vector, w *Vector) *Vector {
	return Make(v.X - w.X, v.Y - w.Y, v.Z - w.Z)
}

func DotProduct(v *Vector, w *Vector) float64 {
	return v.X * w.X + v.Y * w.Y + v.Z * w.Z
}

func Orthogonal(v *Vector, w *Vector) bool {
	return fAlmostEqual(DotProduct(v, w), 0)
}

func SameDirection(v *Vector, w *Vector) bool {
	return relationship(v, w) == 2
}

func OppositeDirection(v *Vector, w *Vector) bool {
	return relationship(v, w) == 1
}

func Parallel(v *Vector, w *Vector) bool {
	return SameDirection(v, w) || OppositeDirection(v, w)
}

func relationship(v *Vector, w *Vector) int {
	vu := Unit(v)
	wu := Unit(w)
	if AlmostEqual(vu, wu) {
		return 2
	} else if AlmostEqual(vu, Negate(wu)) {
		return 1
	}
	return 0 
}

func fAlmostEqual(x float64, y float64) bool {
	return math.Fabs(y - x) < 0.000001
}
