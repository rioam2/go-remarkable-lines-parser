package goremarkablelinesparser

import (
	"encoding/binary"
	"io"
	"math"
)

func readFloat32(r io.Reader) float32 {
	bytes := make([]byte, 4)
	io.ReadFull(r, bytes)
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func readUint32(r io.Reader) uint32 {
	bytes := make([]byte, 4)
	io.ReadFull(r, bytes)
	bits := binary.LittleEndian.Uint32(bytes)
	return bits
}
