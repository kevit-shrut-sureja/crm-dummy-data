package main

import (
	"fmt"
	"math/rand/v2"
)

func randomPhoneNumber() string {
	num := rand.Int64N(9000000000) + 1000000000 // Generate a 10-digit number
	return fmt.Sprintf("%010d", num)            // Format as 10-digit string
}

func getRandomRealPhoneNumber() string {
	s := probArray(0.2, []string{"8160460050", "9909138575"}, true)
	if len(s) == 0 {
		return ""
	}
	return s[0]
}
