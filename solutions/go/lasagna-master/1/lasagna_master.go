package lasagnamaster

func PreparationTime(layers []string, average int) int {
	time := average
	if average == 0 {
		time = 2
	}

	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	noddles, sauce := 0, 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noddles += 50
		}

		if layer == "sauce" {
			sauce += 0.2
		}
	}

	return noddles, sauce
}

func AddSecretIngredient(friend []string, list []string) {
	index := make(map[string]bool)

	for _, ing := range list {
		index[ing] = true
	}

	for _, ing := range friend {
		_, exists := index[ing]

		if !exists {
			list[len(list)-1] = ing
		}
	}
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	needed := make([]float64, len(quantities))
	size := float64(portions) / 2

	for index, quantity := range quantities {
		needed[index] = size * quantity
	}

	return needed
}

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
