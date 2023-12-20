package player

import (
	"go-gusanos/gameData"
	"go-gusanos/util"
	"image"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	frameWidth  int = 10
	frameHeight int = 10
	frameCount  int = 4
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
	Skin, Mask                              *ebiten.Image
	Crosshair                               *ebiten.Image
	CurrentFirecone                         *ebiten.Image
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
	var frameGutter int = 0

	if frame > 0 {
		frameGutter = 1
	}

	// w.Skin.Bounds().Size()

	// frame is the frame number of skin/mask sprite which should be rendered
	sx := frame * frameWidth
	sy := 0

	ebitenutil.DebugPrintAt(screen, strconv.Itoa(sx), 10, 10)
	ebitenutil.DebugPrintAt(screen, strconv.Itoa(sy), 10, 30)
	ebitenutil.DebugPrintAt(screen, strconv.Itoa(frame), 10, 50)
	// ebitenutil.DebugPrintAt(screen, strconv.Itoa(frameGutter), 10, 70)

	skin, _ := util.PrepareSpriteMap(w.Skin)
	screen.DrawImage(skin.SubImage(image.Rect(sx, sy, sx+frameWidth-frameGutter, sy+frameHeight)).(*ebiten.Image), &ebiten.DrawImageOptions{})

	crosshair := util.CutOffset(w.Crosshair, 1, 1)
	crosshairBounds := crosshair.Bounds().Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(frameWidth)+10, float64(frameHeight/2-crosshairBounds.Y/2))

	screen.DrawImage(crosshair, op)
}

func New(assets gameData.Sprites) Worm {
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
	worm.Crosshair = assets["crosshair.png"] // load image here CHECK
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
