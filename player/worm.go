package player

import (
	"go-gusanos/weapon"
	"image/color"

	"github.com/hajimehoshi/ebiten"
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
	Aim, AimSpeed, AimRecoilSpeed           int64
	CurrentFrame                            uint64
	Direction                               int // enum or bool here?
	Skin, Mask                              ebiten.Image
	Crosshair                               ebiten.Image
	CurrentFirecone                         ebiten.Image
	FireconeTime                            int64
	View                                    ebiten.Image
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
	RenderFlip                              func(where ebiten.Image, frame int64, x, y int64)
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

func (w Worm) Render(where ebiten.Image, frame int64, x, y int64) {
	// frame is the frame number of skin/mask sprite which should be rendered CHECK

	// var R, G, B int64
	// var h1, s1, v1, h, s, v int64
	// var g, c color.Color

	width, height := w.Skin.Size()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			// void worm::render(BITMAP *where, int frame, int _x, int _y) player.cpp
		}
	}
}

func New(weaponsList weapon.WeaponsList) Worm {
	worm := Worm{}

	worm.Name = "Player"
	worm.X = 80 * 1000
	worm.Y = 40 * 1000
	worm.XSpeed = 0
	worm.YSpeed = 0
	worm.Aim = 64000
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
	worm.Crosshair = ebiten.Image{} // load image here CHECK
	worm.Active = false
	worm.Flag = false
	worm.IsLocal = false
	worm.SelectingWeapons = false
	worm.RopeState = 0
	worm.RopeX = 1
	worm.RopeXSpeed = 1
	worm.RopeY = 1
	worm.RopeYSpeed = 1
	worm.CurrentFirecone = ebiten.Image{}
	worm.FireconeTime = 0
	// worm.AimAcceleration = 100
	// worm.AimFriction = 50
	// worm.AimMaxSpeed = 1200
	worm.Skin = ebiten.Image{}
	worm.Mask = ebiten.Image{}
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
