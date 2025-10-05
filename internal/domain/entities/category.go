package entities

import (
	"math/rand/v2"
)

type Category int

const (
	CategoryAnimals Category = iota + 1
	CategoryFruitsVegetables
	CategoryCountries
	CategoryRandom
	CategoryUnknown
)

var playableCategories = []Category{CategoryAnimals, CategoryFruitsVegetables, CategoryCountries}

// AllCategories() returns a list of categories the user can choose.
func AllCategories() []Category {
	return []Category{CategoryAnimals, CategoryFruitsVegetables, CategoryCountries, CategoryRandom}
}

func RandomCategory() Category {
	return playableCategories[rand.IntN(len(playableCategories))]
}

func (c Category) String() string {
	switch c {
	case CategoryAnimals:
		return "Animals"
	case CategoryFruitsVegetables:
		return "Fruits & Vegetables"
	case CategoryCountries:
		return "Countries"
	case CategoryRandom:
		return "Random"
	}
	return "Invalid category"
}
