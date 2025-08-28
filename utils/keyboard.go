package utils

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type KeyboardLayout struct {
	Modifier     int32
	Plus         int32
	Minus        int32
	PlayerLeft   int32
	PlayerRight  int32
	PlayerTop    int32
	PlayerBottom int32
}

func InitKeyboardLayoutAzerty() KeyboardLayout {
	return KeyboardLayout{
		Modifier:     rl.KeyLeftSuper,
		Plus:         rl.KeySlash,
		Minus:        rl.KeyEqual,
		PlayerLeft:   rl.KeyA,
		PlayerRight:  rl.KeyD,
		PlayerTop:    rl.KeyW,
		PlayerBottom: rl.KeyS,
	}
}

func InitKeyboardLayoutQwerty() KeyboardLayout {
	return KeyboardLayout{
		PlayerLeft:   rl.KeyQ,
		PlayerRight:  rl.KeyD,
		PlayerTop:    rl.KeyZ,
		PlayerBottom: rl.KeyS,
	}
}
