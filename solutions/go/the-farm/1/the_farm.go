package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cows int
}

func (e *InvalidCowsError) Error() string {
	message := "there are no negative cows"
	if e.cows == 0 {
		message = "no cows don't need food"
	}

	return fmt.Sprintf("%d cows are invalid: %s", e.cows, message)
}

func DivideFood(calculator FodderCalculator, cows int) (float64, error) {
	amount, err := calculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (amount * factor) / float64(cows), nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	amount, err := DivideFood(calculator, cows)
	if err != nil {
		return 0, err
	}

	return amount, nil
}

func ValidateNumberOfCows(cows int) error {
	if cows <= 0 {
		return &InvalidCowsError{cows}
	}

	return nil
}
