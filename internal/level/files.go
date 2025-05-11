package level

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (l *Level) Load(path string) error {
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

	l.Width = int32(len(l.Grid[0]))
	l.Height = int32(len(l.Grid))
	l.Current = path

	var counter int32 = 0
	for i := range l.Grid {
		for _, cell := range l.Grid[i] {
			if cell == Point || cell == Power {
				counter++
			}
		}
	}
	l.Points = counter

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
	file, err := os.Create(path)
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

func (l *Level) CreateFile(folder string) {
	filename, _ := createUniqueFile(folder + "/level.json")
	l.LoadDefault()
	l.Save(filename)
}

func createUniqueFile(baseName string) (string, error) {
	ext := filepath.Ext(baseName)
	name := strings.TrimSuffix(baseName, ext)
	counter := 1
	filename := fmt.Sprintf("%s%d%s", name, counter, ext)

	for {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			break
		}
		filename = fmt.Sprintf("%s%d%s", name, counter, ext)
		counter++
	}

	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return filename, nil
}
