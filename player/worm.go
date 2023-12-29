package player

import (
	"go-gusanos/gameData"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	crosshairDistance = 20
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
	CrossR                                  int64 // ???
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
	RenderFlip                              func(where *ebiten.Image, frame int64, x, y int64)
}

func (w Worm) SendMessage(message string) {
	// implementation of sending message to chat
}

func (w Worm) SendShootEvents() {
	// implementation of sending shoot event to network stream
	// void worm::shooteventsend()
}

func (w Worm) SendDeathEvents() {
	// implementation of sending death event to network stream
	// void worm::deatheventsend()
}

func (w Worm) CheckEvents() {
	// implementation of updating events affecting player
	// void worm::checkevents()
}

func (w Worm) Render(screen *ebiten.Image, frame int) {
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
		float64(w.X)+crosshairDistance*math.Cos(float64(w.Aim)),
		float64(w.Y)+crosshairDistance*math.Sin(float64(w.Aim)))

	screen.DrawImage(w.Crosshair.Image, op)
}

func New(assets gameData.Sprites) Worm {
	worm := Worm{}

	worm.Name = "Player"
	worm.X = 80
	worm.Y = 40
	worm.XSpeed = 0
	worm.YSpeed = 0
	worm.Aim = 0
	worm.Direction = 1
	worm.CrossR = 20
	worm.Health = 1000
	worm.Deaths = 0
	worm.CurrentWeapon = 0
	worm.AimSpeed = 0
	worm.AimRecoilSpeed = 0
	worm.Color = color.RGBA{100, 100, 220, 1}

	// for _, weapon := range worm.Weapon {
	// 	weapon.WeaponIndex = 0
	// 	weapon.ShootTime = 0
	// 	weapon.Ammo = weaponsList.Number[weapon.WeaponIndex].Ammo
	// 	weapon.Reloading = false
	// }

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

	return worm
}
