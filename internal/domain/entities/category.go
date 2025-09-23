package entities

import (
	"math/rand/v2"
	"slices"
)

type Category int

const (
	CategoryAnimals Category = iota + 1
	CategoryFruitsVegetables
	CategoryCountries
	CategoryRandom
	CategoryUnknown
)

var (
	allCategories      = []Category{CategoryAnimals, CategoryFruitsVegetables, CategoryCountries, CategoryRandom, CategoryUnknown} // all existing categories
	playableCategories = []Category{CategoryAnimals, CategoryFruitsVegetables, CategoryCountries}
)

func (c Category) IsValid() bool {
	return slices.Contains(allCategories, c)
}

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
