package gameData

type GameDataRepository struct {
	Sprites Sprites
}

func New(modName string) GameDataRepository {
	repository := GameDataRepository{
		Sprites: LoadSprites(modName),
	}

	return repository
}
