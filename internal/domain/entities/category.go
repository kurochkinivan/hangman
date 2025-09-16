package entities

type Category int 

const (
	CategoryAnimals Category = iota + 1
	CategoryFruitsVegetables
	CategoryCountries
)

var categoryNames = map[Category]string{
	CategoryAnimals:          "Animals",
	CategoryFruitsVegetables: "Fruits & Vegetables",
	CategoryCountries:        "Countries",
}

func (c Category) String() string {
	if name, ok := categoryNames[c]; ok {
		return name
	}
	return "Unknown"
}

func (c Category) IsValid() bool {
	_, ok := categoryNames[c]
	return ok 
}

func AllCategories() []Category {
	return []Category{CategoryAnimals, CategoryFruitsVegetables, CategoryCountries}
}


