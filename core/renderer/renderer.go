package renderer

import (
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EntityData struct {
	Layer    int
	Position rl.Vector2
	Size     rl.Vector2
}

type RenderableObject interface {
	GetEntityData() EntityData

	AsRect() rl.Rectangle

	Render()
}

type Renderer struct {
	currentId int
	objects   map[int]RenderableObject
}

func InitRenderer() Renderer {
	return Renderer{currentId: 0, objects: map[int]RenderableObject{}}
}

func (renderer *Renderer) AddObject(object RenderableObject) int {
	id := renderer.currentId
	renderer.objects[id] = object
	renderer.currentId += 1

	return id
}

func (renderer *Renderer) UpdateObject(id int, object RenderableObject) {
	renderer.objects[id] = object
}

func (renderer Renderer) Render() {
	for _, object := range renderer.sortObjectsForRendering() {
		object.Render()
	}
}

func (renderer Renderer) sortObjectsForRendering() []RenderableObject {
	sorted := []RenderableObject{}

	for _, object := range renderer.objects {
		sorted = append(sorted, object)
	}

	slices.SortStableFunc(sorted, func(a, b RenderableObject) int {
		aLayer := a.GetEntityData().Layer
		bLayer := b.GetEntityData().Layer

		if aLayer < bLayer {
			return -1
		} else if aLayer > bLayer {
			return 1
		} else {
			return 0
		}
	})

	return sorted
}
