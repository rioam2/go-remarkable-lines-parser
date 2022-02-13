package goremarkablelinesparser

import (
	"io"
)

type Layer struct {
	NumLines uint32
	Lines    []*Line
}

func (layer *Layer) Read(r io.Reader) error {
	var err error
	layer.NumLines = readUint32(r)
	for lineNum := 0; lineNum < int(layer.NumLines); lineNum++ {
		line := &Line{}
		err = line.Read(r)
		if err != nil {
			return err
		}
		layer.Lines = append(layer.Lines, line)
	}
	return nil
}
