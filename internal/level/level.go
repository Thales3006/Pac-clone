package level

import (
	"encoding/json"
	"io"
	"os"

	"github.com/dominikbraun/graph"
)

type Level struct {
	Grid    [][]Side                    `json:"Grid"`
	Width   uint8                       `json:"-"`
	Height  uint8                       `json:"-"`
	Current string                      `json:"-"`
	Graph   graph.Graph[[2]uint8, Cell] `json:"-"`
}

type Side uint8

const (
	Empty Side = iota
	Wall
	Door
	Point
)

func NewLevel() *Level {
	return &Level{
		Grid:    nil,
		Width:   0,
		Height:  0,
		Current: "",
	}
}

func (l *Level) Load(path string) error {
	file, err := os.Open("levels/" + path)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, l)
	if err != nil {
		return err
	}

	l.Width = uint8(len(l.Grid[0]))
	l.Height = uint8(len(l.Grid))
	l.Current = path

	l.generateGraph()

	return nil
}

func (l *Level) Unload() {
	l.Grid = nil
	l.Width = 0
	l.Height = 0
	l.Current = ""
	l.Graph = nil
}

func (l *Level) Save(path string) error {
	file, err := os.Create("levels/" + path)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	return err
}
