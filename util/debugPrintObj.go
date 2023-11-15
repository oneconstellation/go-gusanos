package util

import (
	"reflect"

	"github.com/hajimehoshi/ebiten"
)

func DebugPrintObj[T any](screen *ebiten.Image, obj T) {
	fields := reflect.VisibleFields(reflect.TypeOf(obj))

	for _, field := range fields {
		reflect.Indirect()
	}
}
