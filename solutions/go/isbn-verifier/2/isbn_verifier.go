package isbnverifier

func IsValidISBN(isbn string) bool {
    sum := 0
	digit := 0

    for i := 0; i < len(isbn); i++ {
		char := isbn[i]

        if char == ' ' || char == '-' {
            continue
        }
        
        if char < '0' || char > '9' {
            if char != 'X' || char == 'X' && digit != 9 {
            	return false
            }        
        }
        
        number := 0       

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
