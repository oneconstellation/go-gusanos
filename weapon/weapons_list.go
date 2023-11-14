package weapon

type WeaponsList struct {
	Start, End   *Weapon
	Number       [100]*Weapon
	WeaponsCount int64
}

func New() WeaponsList {
	return WeaponsList{}
}
