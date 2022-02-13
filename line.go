package goremarkablelinesparser

import (
	"io"
)

type Line struct {
	Type      uint32
	Color     uint32
	Size      float32
	NumPoints uint32
	Points    []*Point
}

func (line *Line) Read(r io.Reader) error {
	var err error
	line.Type = readUint32(r)
	line.Color = readUint32(r)
	_ = readUint32(r) // padding
	line.Size = readFloat32(r)
	_ = readUint32(r) // padding
	line.NumPoints = readUint32(r)
	for pointNum := 0; pointNum < int(line.NumPoints); pointNum++ {
		point := &Point{}
		err = point.Read(r)
		if err != nil {
			return err
		}
		line.Points = append(line.Points, point)
	}
	return nil
}
