package weapon

import (
	"go-gusanos/particle"

	"github.com/hajimehoshi/ebiten"
)

type Weapon struct {
	ShootSpeed                   int64
	ShootNumber                  int64
	Distribution                 int64
	ShootTimes                   int64
	ShootSpeedRand               int64
	AimRecoil                    int64
	Recoil                       int64
	AffectedMotion               int64
	FireconeTimeout              int64
	Ammo                         int64
	ReloadTime                   int64
	LasersightIntensity          int64
	LasersightFade               int64
	StartDelay                   int64
	Autofire                     bool
	ShootObject, CreateOnRelease particle.ParticleType
	Firecone                     ebiten.Image
	Name, FileName               string
	Next, Previous               *Weapon
	// ShootSound, ReloadSound, NoAmmoSound, StartSound
}
