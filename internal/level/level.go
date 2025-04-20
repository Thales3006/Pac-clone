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
	Loaded   bool
	Unlocked uint8
	Current  string
}

type Side uint8

const (
	Empty Side = iota
	Wall
	Door
)

func NewLevel() *Level {
	return &Level{
		Grid:     nil,
		Width:    0,
		Height:   0,
		Unlocked: 1,
		Current:  "",
		Loaded:   false,
	}
}

func (l *Level) Load() error {
	file, err := os.Open("levels/" + l.Current + ".json")
	if err != nil {
		l.Loaded = false
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		l.Loaded = false
		return err
	}
	err = json.Unmarshal(bytes, l)
	if err != nil {
		l.Loaded = false
		return err
	}

	l.Loaded = true
	return err
}

func (l *Level) Unload(path string) {
	l.Grid = nil
	l.Width = 0
	l.Height = 0
	l.Unlocked = 1
	l.Current = ""
	l.Loaded = false
}
