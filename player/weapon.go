package player

type Weapon struct {
	WeaponIndex                             int64
	Ammo, ShootTime, ReloadTime, StartDelay int64
	Reloading                               bool
}
