package allergies

var allergens = []string{
	"eggs",         // 1   = 1 << 0
	"peanuts",      // 2   = 1 << 1
	"shellfish",    // 4   = 1 << 2
	"strawberries", // 8   = 1 << 3
	"tomatoes",     // 16  = 1 << 4
	"chocolate",    // 32  = 1 << 5
	"pollen",       // 64  = 1 << 6
	"cats",         // 128 = 1 << 7
}

func Allergies(allergies uint) []string {
	var result []string
	for i, name := range allergens {
		if allergies&(1<<i) != 0 {
			result = append(result, name)
		}
	}
	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	for i, name := range allergens {
		if name == allergen {
			return allergies&(1<<i) != 0
		}
	}
	return false
}
