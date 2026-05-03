package twofer

import "fmt"

func ShareWith(name string) string {
	with := name
    
    if name == "" {
        with = "you"
    }
    
	return fmt.Sprintf("One for %s, one for me.", with)
}
