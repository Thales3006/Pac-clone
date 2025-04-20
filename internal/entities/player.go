package entities

type Player struct {
	Health uint8
	X      float32
	Y      float32
	Height float32
	Width  float32
	Vel    float32

	LevelUnlocked int16
	LevelCurrent  int16
}

func NewPlayer() *Player {
	return &Player{
		Health: 5,
		X:      5,
		Y:      10,
		Height: 1,
		Width:  1,
		Vel:    0.1,
	}
}
