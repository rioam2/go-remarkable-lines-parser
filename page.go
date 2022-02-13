package goremarkablelinesparser

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

const RM_LINES_FILE_MAGIC = "reMarkable .lines file, version="

type Page struct {
	NumLayers uint32
	Layers    []*Layer
}

func (page *Page) Read(r io.Reader) error {
	var err error
	page.NumLayers = readUint32(r)
	for layerNum := 0; layerNum < int(page.NumLayers); layerNum++ {
		layer := &Layer{}
		err = layer.Read(r)
		if err != nil {
			return err
		}
		page.Layers = append(page.Layers, layer)
	}
	return nil
}

func FromReader(r io.Reader) (*Page, error) {
	magicBytes := make([]byte, 32)
	io.ReadFull(r, magicBytes)
	magic := string(magicBytes)
	if magic != RM_LINES_FILE_MAGIC {
		return nil, fmt.Errorf("invalid file header")
	}
	versionBytes := make([]byte, 11)
	io.ReadFull(r, versionBytes)
	version, err := strconv.ParseUint(strings.TrimSpace(string(versionBytes)), 10, 64)
	if err != nil {
		return nil, err
	}
	if version != 5 {
		return nil, fmt.Errorf("parsing for version %d files is not supported", version)
	}
	page := &Page{}
	err = page.Read(r)
	if err != nil {
		return nil, err
	}
	return page, nil
}
