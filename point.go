package goremarkablelinesparser

import (
	"io"
)

type Point struct {
	X         float32
	Y         float32
	Speed     float32
	Direction float32
	Width     float32
	Pressure  float32
}

func (point *Point) Read(r io.Reader) error {
	point.X = readFloat32(r)
	point.Y = readFloat32(r)
	point.Speed = readFloat32(r)
	point.Direction = readFloat32(r)
	point.Width = readFloat32(r)
	point.Pressure = readFloat32(r)
	return nil
}
