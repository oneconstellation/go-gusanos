package player

import (
	"go-gusanos/gameData"
	"go-gusanos/util"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Worm struct {
	Name                                    string
	X, Y                                    int64
	XSpeed, YSpeed                          int64
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
	CurrentFrame                            uint64
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

// temporary, hardcoded implementation
func (w *Worm) Update(keys []ebiten.Key) {
	w.Keys.Left = util.Contains(keys, ebiten.KeyA)
	if util.Contains(keys, ebiten.KeyA) {
		w.moveLeft()
	}
	w.Keys.Right = util.Contains(keys, ebiten.KeyD)
	if util.Contains(keys, ebiten.KeyD) {
		w.moveRight()
	}
	w.Keys.Up = util.Contains(keys, ebiten.KeyW)
	if util.Contains(keys, ebiten.KeyW) {
		w.aimUp()
	}
	w.Keys.Down = util.Contains(keys, ebiten.KeyS)
	if util.Contains(keys, ebiten.KeyS) {
		w.aimDown()
	}
	w.Keys.Fire = util.Contains(keys, ebiten.KeyF)
	w.Keys.Jump = util.Contains(keys, ebiten.KeyG)
	w.Keys.Change = util.Contains(keys, ebiten.KeyH)
}

func (w *Worm) aimUp() {
	w.Aim -= w.AimSpeed
}

func (w *Worm) aimDown() {
	w.Aim += w.AimSpeed
}

func (w *Worm) moveLeft() {
	if w.Direction != 0 {
		w.Direction = 0
	}
}

func (w *Worm) moveRight() {
	if w.Direction != 1 {
		w.Direction = 1
	}
}

func (w *Worm) Render(screen *ebiten.Image, frame int) {
	// render player skin
	skin, op := w.Skin.GetSubSprite(frame, 0, w.Direction == 0, false)
	op.GeoM.Translate(float64(w.X), float64(w.Y))
	screen.DrawImage(skin, op)

	// render player mask - TODO change color and blend with skin sprite
	mask, op := w.Mask.GetSubSprite(frame, 0, w.Direction == 0, false)
	op.GeoM.Translate(float64(w.X), float64(w.Y))
	screen.DrawImage(mask, op)

	// render crosshair
	crosshairSize := w.Crosshair.Image.Bounds().Size()
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(crosshairSize.X/2), -float64(crosshairSize.Y/2))

	op.GeoM.Translate(
		float64(w.X)+float64(w.CrossR)*math.Cos(float64(w.Aim)),
		float64(w.Y)+float64(w.CrossR)*math.Sin(float64(w.Aim)))

	screen.DrawImage(w.Crosshair.Image, op)
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

	worm.CurrentFrame = 2700
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
	// worm.AimAcceleration = 100
	// worm.AimFriction = 50
	// worm.AimMaxSpeed = 1200
	worm.Skin = assets["skin.png"]
	worm.Mask = assets["skin-mask.png"]
	worm.Keys = Keys{
		Up:     false,
		Down:   false,
		Right:  false,
		Left:   false,
		Jump:   false,
		Fire:   false,
		Change: false,
	}

	// here should go the replication of values to network stream
	// with safety checks

	return &worm
}
