package gameData

type GameDataRepository struct {
	Sprites Sprites
	Maps    Maps
}

func New(modName string) GameDataRepository {
	repository := GameDataRepository{
		Sprites: LoadSprites(modName),
		Maps:    LoadMaps(modName),
	}

	return repository
}
