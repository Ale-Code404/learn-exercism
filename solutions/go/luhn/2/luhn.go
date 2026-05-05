package luhn

func Valid(id string) bool {
	if len(id) <= 1 {
        return false
    }

    sum := 0
    processed := 0

    for i := len(id) - 1; i >= 0; i-- {
		number := id[i]
		if (number == ' ') {
            continue
        }

        if number < '0' || number > '9' {
            return false
        }
                
        var converted int = int(number - '0')        

        processed++

        if processed % 2 == 0 {
			double := converted * 2
            sum += double
            
            if double > 9 {
            	sum -= 9            
            }
        } else {
            sum += converted
        }
    }

    if processed <= 1 {
        return false
    }

    return sum % 10 == 0
}
