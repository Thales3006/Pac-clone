package level

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Level struct {
	Grid     [][]Side
	Width    uint8
	Height   uint8
	Unlocked uint8
	Current  uint8
}

type Side uint8

const (
	Empty Side = iota
	Wall
	Door
)

func LoadLevel(path string) (*Level, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var level Level
	err = json.Unmarshal(bytes, &level)
	if err != nil {
		return nil, err
	}

	return &level, nil
}

func NewLevel() *Level {
	grid := make([][]Side, 20)

	for i := range grid {
		grid[i] = make([]Side, 20)
		for j := range grid[i] {
			//totally arbitrary setup
			if i == j || i == 0 || j == 0 {
				grid[i][j] = Wall
			} else {
				grid[i][j] = Empty
			}
		}
	}

	return &Level{
		Grid:     grid,
		Width:    uint8(len(grid[0])),
		Height:   uint8(len(grid)),
		Unlocked: 1,
		Current:  0,
	}
}
