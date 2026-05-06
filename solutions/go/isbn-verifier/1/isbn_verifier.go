package isbnverifier

func IsValidISBN(isbn string) bool {
    sum := 0
	digit := 0

    for i := 0; i < len(isbn); i++ {
		char := isbn[i]
        number := 0

        if char < '0' || char > '9' {		
			if char != ' ' && char != '-' && digit != 9 {
                return false
            }

            if char != 'X' {
            	continue
            }
        }

        if char == 'X' {
            number = 10
        } else {
        	number = int(char - '0')
        }    

        digit++
        sum += number * (11 - digit)
    }

    if digit != 10 {
        return false
    }

    return sum % 11 == 0
}
