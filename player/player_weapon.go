package player

type PlayerWeapon struct {
	WeaponIndex                             int64
	Ammo, ShootTime, ReloadTime, StartDelay int64
	Reloading                               bool
}
