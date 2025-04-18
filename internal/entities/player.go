package entities

type Player struct {
	Health uint8
	X      float32
	Y      float32
	Height float32
	Width  float32

	LevelUnlocked int16
	LevelCurrent  int16
}

func NewPlayer() *Player {
	return &Player{
		Health:        5,
		X:             0,
		Y:             0,
		Height:        0.7,
		Width:         0.7,
		LevelUnlocked: 1,
		LevelCurrent:  0,
	}
}
