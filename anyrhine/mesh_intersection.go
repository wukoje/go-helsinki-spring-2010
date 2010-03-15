/*
 *	Triangle mesh intersection with a ray
 *
 *	test with:
 *		8.mesh_intersection < mesh_instersection.t1
 *		8.mesh_intersection < mesh_instersection.t2
 *
 *	todo: extend to a raytracer.
 *
 *	Ray-triangle intersection code adapted from [mt97]
 *	check there for explanation.
 *
 *	[mt97]	Fast, minimum storage ray-triangle intersection.
 *		Tomas Möller and Ben Trumbore.
 *		Journal of Graphics Tools, 2(1):21--28, 1997.
 */

package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

type Vtx [3]float
type Tri [3]*Vtx

const Epsilon = 0.000001

func dot(v1 *Vtx, v2 *Vtx) float {
	a, b := (*[3]float)(v1), (*[3]float)(v2)
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func cross(v1 *Vtx, v2 *Vtx) Vtx {
	a, b := (*[3]float)(v1), (*[3]float)(v2)
	return Vtx{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}
}

func sub(v1 *Vtx, v2 *Vtx) Vtx {
	a, b := (*[3]float)(v1), (*[3]float)(v2)
	return Vtx{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func isect(orig *Vtx, dir *Vtx, triv *Tri) (Vtx, bool) {
	tri := [3]*Vtx(*triv)
	e1, e2 := sub(tri[1], tri[0]), sub(tri[2], tri[0])
	pvec := cross(dir, &e2)
	det := dot(&e1, &pvec)
	if det > -Epsilon && det < Epsilon {
		return Vtx{}, false
	}
	idet := 1.0 / det
	tvec := sub(orig, tri[0])
	u := dot(&tvec, &pvec) * idet
	if u < 0.0 || u > 1.0 {
		return Vtx{}, false
	}
	qvec := cross(&tvec, &e1)
	v := dot(dir, &qvec) * idet
	if v < 0.0 || u+v > 1.0 {
		return Vtx{}, false
	}
	t := dot(&e2, &qvec) * idet
	return Vtx{t, u, v}, true
}

var vtxmap = map[string]*Vtx{}
var trimap = map[string]*Tri{}
var lampmap = map[string]*Vtx{}
var cam *Vtx
var lookat *Vtx

func hasvtx(tok []string) bool {
	rval := true
	for i := 0; i < len(tok); i++ {
		if _, ok := vtxmap[tok[i]]; !ok {
			fmt.Printf("error: vtx %s undefined\n", tok[i])
			rval = false
		}
	}
	return rval
}

func parse() {
	r := bufio.NewReader(os.Stdin)
	for {
		if s, err := r.ReadString('\n'); err == nil {
			tok := strings.Fields(s)
			if len(tok) < 1 {
				goto usage
			}
			switch tok[0] {
			case "vtx":
				if len(tok) != 5 {
					goto usage
				}
				a, e1 := strconv.Atof(tok[2])
				b, e2 := strconv.Atof(tok[3])
				c, e3 := strconv.Atof(tok[4])
				if e1 != nil || e2 != nil || e3 != nil {
					goto usage
				}
				vtxmap[tok[1]] = &Vtx{a, b, c}
			case "tri":
				if len(tok) != 5 || !hasvtx(tok[2:5]) {
					goto usage
				}
				trimap[tok[1]] = &Tri{vtxmap[tok[2]], vtxmap[tok[3]], vtxmap[tok[4]]}
			case "cam":
				if len(tok) != 3 || !hasvtx(tok[1:3]) {
					goto usage
				}
				cam = vtxmap[tok[1]]
				lookat = vtxmap[tok[2]]
			case "lamp":
				if len(tok) != 3 || !hasvtx(tok[2:3]) {
					goto usage
				}
				lampmap[tok[1]] = vtxmap[tok[2]]
			case "q":
				goto quit
			default:
			usage:
				fmt.Print("usage: vtx <vtx-id> x y z\n")
				fmt.Print("       tri <tri-id> <vtx1> <vtx2> <vtx3>\n")
				fmt.Print("       lamp <lamp-id> <pos-vtx>\n")
				fmt.Print("       cam <pos-vtx> <lookat-vtx>\n")
				fmt.Print("       quit\n")
			}
		} else {
		quit:
			break
		}
	}
}

func main() {
	parse()
	for _, tri := range trimap {
		if _, ok := isect(cam, lookat, tri); ok {
			fmt.Printf("intersects\n")
		} else {
			fmt.Printf("does not intersect\n")
		}
	}
}
