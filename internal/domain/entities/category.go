package entities

import "math/rand/v2"

type Category int

const (
	CategoryAnimals Category = iota + 1
	CategoryFruitsVegetables
	CategoryCountries
	CategoryRandom
	CategoryUnknown
)

var categoryNames = map[Category]string{
	CategoryAnimals:          "Animals",
	CategoryFruitsVegetables: "Fruits & Vegetables",
	CategoryCountries:        "Countries",
	CategoryRandom:           "Random",
}

func (c Category) String() string {
	if name, ok := categoryNames[c]; ok {
		return name
	}
	return "Invalid category"
}

func (c Category) IsValid() bool {
	_, ok := categoryNames[c]
	return ok
}

func AllCategories() []Category {
	categories := make([]Category, 0, len(categoryNames))

	for category := range categoryNames {
		categories = append(categories, category)
	}

	return categories
}

func RandomCategory() Category {
	categories := []Category{CategoryAnimals, CategoryFruitsVegetables, CategoryCountries}
	return categories[rand.IntN(len(categories))]
}
