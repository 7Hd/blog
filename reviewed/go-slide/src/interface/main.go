package main

import "fmt"

type IPoint interface {
	X() float64
	Y() float64
	SetX(float64)
	SetY(float64)
}

type IPoint3D interface {
	IPoint
	Z() float64
	SetZ(float64)
}

type Point struct {
	x float64
	y float64
}

// START OMIT
func (p *Point) X() float64     { return p.x }
func (p *Point) Y() float64     { return p.y }
func (p *Point) SetX(x float64) { p.x = x }
func (p *Point) SetY(y float64) { p.y = y }

// END OMIT

type Point3D struct {
	Point
	z float64
}

func (p *Point3D) Z() float64     { return p.z }
func (p *Point3D) SetZ(z float64) { p.z = z }

func main() {
	var p IPoint
	p = &Point{}
	p.SetX(4)
	p.SetY(3)

	var p3d IPoint3D = &Point3D{}
	p3d.SetX(4)
	p3d.SetY(3)
	p3d.SetZ(2)

	fmt.Printf("%#v\n", p)
	fmt.Printf("%#v\n", p3d)
}
