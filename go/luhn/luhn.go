package luhn

import (
	"log"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 {
		return false
	}

	sum, double := 0, len(id)%2 == 0
	for _, token := range id {
		digit, err := strconv.Atoi(string(token))
		if err != nil {
			log.Println(err)
			return false
		}

		if double {
			digit = digit * 2

			if digit > 9 {
				digit -= 9
			}
		}

		double = !double
		sum += digit
	}

	return sum%10 == 0
}
