package player

import (
	"fmt"
	"go-gusanos/gameData"
	"image/color"
	"math"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	animationFrames int     = 4
	animationSpeed  float64 = 0.1
	wormSpeed       float64 = 0.5
	wormJumpForce   float64 = 25.0
	gravity         float64 = 1.0 // TODO move to game
)

type Worm struct {
	Name                                    string
	X, Y                                    float64
	XSpeed, YSpeed                          float64
	RopeX, RopeY                            int64
	RopeXSpeed, RopeYSpeed                  int64
	RopeLength                              int64
	RopeState                               int // enum here?
	XView, YView                            int64
	CrosshairX, CrosshairY                  int64
	Health                                  int64
	Deaths                                  int64
	Air                                     int64
	CrossR                                  int64
	Aim, AimSpeed, AimRecoilSpeed           float64
	CurrentFrame                            int
	Direction                               int // 0 - left, 1 - right
	Skin, Mask                              gameData.Sprite
	Crosshair                               gameData.Sprite
	CurrentFirecone                         gameData.Sprite
	FireconeTime                            int64
	View                                    *ebiten.Image
	Weapon                                  [5]Weapon
	CurrentWeapon                           int64
	Active, IsLocal                         bool
	Flag, Flag2, Flag3, FlagLeft, FlagRight bool
	RopeFlag, SelectingWeapons, Firing      bool
	ShootRope                               func()
	DestroyRope                             func()
	ApplyRopeForce                          func()
	Keys                                    Keys
	Color                                   color.Color
}

type Function func()

func (w *Worm) Update(keys []ebiten.Key) {
	w.updateAnimation()
	w.listenForKeys(keys)
	w.updateMovement()
}

func (w *Worm) listenForKeys(keys []ebiten.Key) {
	keymap := map[ebiten.Key]Function{
		ebiten.KeyA: w.moveLeft,
		ebiten.KeyD: w.moveRight,
		ebiten.KeyW: w.aimUp,
		ebiten.KeyS: w.aimDown,
		ebiten.KeyF: w.fire,
		ebiten.KeyG: w.jump,
		ebiten.KeyH: w.weaponChangeMode,
	}

	for key, action := range keymap {
		if slices.Contains(keys, key) {
			action()
		}
	}
}

func (w *Worm) updateAnimation() {
	if w.CurrentFrame < animationFrames*10 {
		w.CurrentFrame++
	} else {
		w.CurrentFrame = 0
	}
}

func (w *Worm) updateMovement() {
	w.X = w.X + w.XSpeed
	w.Y = w.Y + w.YSpeed

	w.XSpeed = 0

	if w.YSpeed > gravity {
		w.YSpeed--
	}
}

func (w *Worm) aimUp() {
	if w.Aim >= -math.Pi/2 {
		w.Aim -= w.AimSpeed
	}
}

func (w *Worm) aimDown() {
	if w.Aim <= math.Pi/2 {
		w.Aim += w.AimSpeed
	}
}

func (w *Worm) fire() {
	fmt.Println(w.Name + ": fire")
}

func (w *Worm) weaponChangeMode() {
	fmt.Println(w.Name + ": weapon change mode")
}

func (w *Worm) moveLeft() {
	if w.Direction != 0 {
		w.Direction = 0
	}
	w.XSpeed = -wormSpeed
}

func (w *Worm) moveRight() {
	if w.Direction != 1 {
		w.Direction = 1
	}
	w.XSpeed = wormSpeed
}

func (w *Worm) jump() {
	if w.YSpeed == gravity {
		w.YSpeed = wormJumpForce
	}
}

func (w *Worm) Render(screen *ebiten.Image) {
	// calculate animationFrame
	var counter float64 = float64(w.CurrentFrame) * animationSpeed
	animationFrame := int(counter) % animationFrames

	// render crosshair
	var angle float64
	crosshairSize := w.Crosshair.Image.Bounds().Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(crosshairSize.X/2), -float64(crosshairSize.Y/2))

	if w.Direction == 0 {
		angle = math.Pi - w.Aim
	} else {
		angle = w.Aim
	}

	op.GeoM.Translate(
		float64(w.X)+float64(w.CrossR)*math.Cos(float64(angle)),
		float64(w.Y)+float64(w.CrossR)*math.Sin(float64(angle)))

	screen.DrawImage(w.Crosshair.Image, op)

	// get aim-dependent frame
	// normalize to positive range from 0 to math.Pi,
	// so add half of pi to aim angle
	// multiply by 2.8... it stays in the 0-8 range
	// without need for modulo
	aimPosRange := w.Aim + math.Pi/2
	aimFrame := int(aimPosRange * 2.8)

	// render player skin
	skin, op := w.Skin.GetSubSprite(int(animationFrame), aimFrame, w.Direction == 0, false)
	op.GeoM.Translate(float64(w.X), float64(w.Y))
	screen.DrawImage(skin, op)

	// render player mask - TODO change color and blend with skin sprite
	mask, op := w.Mask.GetSubSprite(int(animationFrame), aimFrame, w.Direction == 0, false)
	op.GeoM.Translate(float64(w.X), float64(w.Y))
	screen.DrawImage(mask, op)
}

func New(assets gameData.Sprites) *Worm {
	worm := Worm{}

	worm.Name = "Player"
	worm.X = 80
	worm.Y = 40
	worm.XSpeed = 0
	worm.YSpeed = 0
	worm.Aim = 0
	worm.AimSpeed = 0.02
	worm.Direction = 1
	worm.CrossR = 20
	worm.Health = 1000
	worm.Deaths = 0
	worm.CurrentWeapon = 0
	worm.AimRecoilSpeed = 0
	worm.Color = color.RGBA{100, 100, 220, 1}
	worm.CurrentFrame = 0
	worm.Crosshair = assets["crosshair.png"]
	worm.Active = false
	worm.Flag = false
	worm.IsLocal = false
	worm.SelectingWeapons = false
	worm.RopeState = 0
	worm.RopeX = 1
	worm.RopeXSpeed = 1
	worm.RopeY = 1
	worm.RopeYSpeed = 1
	worm.CurrentFirecone = assets["firecone.png"]
	worm.FireconeTime = 0
	worm.Skin = assets["skin.png"]
	worm.Mask = assets["skin-mask.png"]
	worm.Keys = Keys{
		Jump:   false,
		Fire:   false,
		Change: false,
	}

	// here should go the replication of values to network stream
	// with safety checks

	return &worm
}
