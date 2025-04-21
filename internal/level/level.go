package level

import (
	"encoding/json"
	"io"
	"os"
)

type Level struct {
	Grid     [][]Side
	Width    uint8
	Height   uint8
	Unlocked uint8
	Current  string
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
		Grid:     nil,
		Width:    0,
		Height:   0,
		Unlocked: 1,
		Current:  "",
	}
}

func (l *Level) Load(path string) error {
	if l.Current != "" {
		return nil
	}
	file, err := os.Open(path)
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

	l.Current = path
	return nil
}

func (l *Level) Unload() {
	l.Grid = nil
	l.Width = 0
	l.Height = 0
	l.Unlocked = 1
	l.Current = ""
}
