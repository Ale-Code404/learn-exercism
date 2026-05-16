package phonenumber

import (
    "fmt"
    "errors"
    "regexp"
)

const NumberLength = 10
const NumberMaxLength = 11

const NumberPrefix = 1
const NumberPrefixUTF8 = 49

const IdentifierMin = 2
const IdentifierMinUTF8 = 50

func Number(phoneNumber string) (string, error) {
	re, err := regexp.Compile(`\D`)
    if err != nil {
        return "", errors.New("Cannot process the number")
    }
    
    number := re.ReplaceAllString(phoneNumber, "")
    length := len(number)
    
    if length > NumberMaxLength {
        return "", errors.New("The max length of number is 11")
    }

    if length == NumberMaxLength && number[0] != NumberPrefixUTF8 {
        return "", errors.New(fmt.Sprintf(
            "A %d phone number length is only valid with a %d prefix",
            NumberMaxLength,
            NumberPrefix,
        ))
    }

    if length < NumberLength {
        return "", errors.New(fmt.Sprintf(
            "The minimun phone number length is %d",
            NumberLength,
        ))
    }

    if length == NumberMaxLength {
    	number = number[1:NumberMaxLength]
    }

    if number[0] < IdentifierMinUTF8 || number[3] < IdentifierMinUTF8 {
        return "", errors.New(fmt.Sprintf(
            "The area and exchange part must start from %d",
            IdentifierMin,
        ))
    }
    
	return number, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
        return "", err
    }

    return number[0:3], nil
}

func Format(phoneNumber string) (string, error) {
    number, err := Number(phoneNumber)
	if err != nil {
        return "", err
    }
    
	area, err := AreaCode(phoneNumber)
    if err != nil {
        return "", err
    }

    return fmt.Sprintf(
        "(%s) %s-%s",
        area,
        number[3:6],
        number[6:10],
	), nil
}
