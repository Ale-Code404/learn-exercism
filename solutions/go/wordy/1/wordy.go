package wordy

import (
	"regexp"
	"strconv"
)

func Answer(question string) (int, bool) {
	if question[0:7] != "What is" {
		return 0, false
	}

	normalized := question[7 : len(question)-1]

	re := regexp.MustCompile(`-?\d+|plus|minus|multiplied by|divided by|\w+`)
	tokens := re.FindAllString(normalized, -1)

	if tokens == nil {
		return 0, false
	}

	length := len(tokens)

	if length%2 == 0 {
		return 0, false
	}

	total := 0
	operator := ""

	for i := range tokens {
		if i == 0 {
			value, err := strconv.Atoi(tokens[i])
			if err != nil {
				return 0, false
			}

			total = value
			continue
		}

		if i%2 != 0 {
			value := tokens[i]

			switch value {
			case "plus":
				operator = "+"
			case "minus":
				operator = "-"
			case "multiplied by":
				operator = "*"
			case "divided by":
				operator = "/"
			default:
				return 0, false
			}
		}

		if i%2 == 0 {
			right, err := strconv.Atoi(tokens[i])
			if err != nil {
				return 0, false
			}

			switch operator {
			case "+":
				total = total + right
			case "-":
				total = total - right
			case "*":
				total = total * right
			case "/":
				total = total / right
			}
		}
	}

	return total, true
}
