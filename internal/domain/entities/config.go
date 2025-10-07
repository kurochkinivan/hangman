package entities

type GameConfig struct {
	level     Level
	category  Category
	wordsRepo WordsRepository
}

type WordsRepository interface {
	RandomWord(level Level, category Category) (Word, error)
}

func NewGameConfig(wordsRepo WordsRepository) *GameConfig {
	return &GameConfig{
		level:     LevelRandom,
		category:  CategoryRandom,
		wordsRepo: wordsRepo,
	}
}

func (gc *GameConfig) SelectWord() (Word, error) {
	level := gc.Level()
	if level == LevelRandom {
		level = RandomLevel()
	}

	category := gc.Category()
	if category == CategoryRandom {
		category = RandomCategory()
	}

	return gc.wordsRepo.RandomWord(level, category)
}

func (gc *GameConfig) Category() Category {
	return gc.category
}

func (gc *GameConfig) SetCategory(category Category) {
	gc.category = category
}

func (gc *GameConfig) Level() Level {
	return gc.level
}

func (gc *GameConfig) SetLevel(level Level) {
	gc.level = level
}
