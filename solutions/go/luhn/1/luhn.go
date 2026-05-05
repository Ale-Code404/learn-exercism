package luhn

import (
    "strconv"
    "fmt"
)

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
                
        converted, error := strconv.Atoi(string(number))
		fmt.Println(converted, error)
        
        if(error != nil) {
            return false
        }

        processed++

        if processed % 2 == 0 {
            sum += converted * 2            
            if converted * 2 > 9 {
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
