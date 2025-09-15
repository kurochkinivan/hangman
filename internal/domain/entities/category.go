package entities

// TODO: Подумать про рандом.
type Category string

const (
	CategoryAnimals          Category = "Animals"
	CategoryFruitsVegetables Category = "Fruits & Vegetables"
	CategoryCountries        Category = "Countries"
)

var categories = []Category{CategoryAnimals, CategoryFruitsVegetables, CategoryCountries}

func AllCategories() []Category {
	return append([]Category(nil), categories...)
}
