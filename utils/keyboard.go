package utils

import rl "github.com/gen2brain/raylib-go/raylib"

type KeyboardLayout struct {
	PlayerLeft   int32
	PlayerRight  int32
	PlayerTop    int32
	PlayerBottom int32
}

func InitKeyboardLayoutAzerty() KeyboardLayout {
	return KeyboardLayout{
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
