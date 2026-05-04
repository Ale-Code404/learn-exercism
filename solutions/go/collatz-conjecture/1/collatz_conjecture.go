package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {    
    var steps int = 0
	var result int = n

    if result <= 0 {
        return 0, errors.New("Only numbers greater than one")
    }
    
	for {
        if result == 1 {
            return steps, nil
        }
        
        if result % 2 == 0 {
            result /= 2
            steps++
        }

        if result == 1 {
            return steps, nil
        }

        if result % 2 != 0 {
            result = result * 3 + 1
            steps++
        }
    }
}
