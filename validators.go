package do_validator

import (
	"log"
	"regexp"
	"strconv"
)

func IsCedula(number string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	number = reg.ReplaceAllString(number, "")

	if len(number) != 11 {
		return false
	}

	return Lunh(number)
}

func IsRnc(number string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	number = reg.ReplaceAllString(number, "")

	if len(number) != 9 {
		return false
	}

	return RncLunh(number)
}

func RncLunh(idnumber string) bool {
	//factors := convertToArrayOfInt(idnumber)
	factors := []int{7, 9, 8, 6, 5, 4, 3, 2}
	var sum int = 0
	for i := len(factors) - 1; i >= 0; i-- {
		sum += factors[i] * int(idnumber[i])
	}
	remaining := sum % 11
	digit := 0

	if remaining == 0 {
		digit = 2
	} else if remaining == 1 {
		digit = 1
	} else {
		digit = 11 - remaining
	}

	lastRncDigit, _ := strconv.Atoi(idnumber[len(idnumber)-1:])
	return digit == lastRncDigit
}

func Lunh(idnumber string) bool {
	digits := len(idnumber)

	var sum uint8 = 0
	isSecond := false
	for i := digits - 1; i >= 0; i-- {
		d := idnumber[i] - '0'
		if isSecond {
			d = d * 2
		}
		sum += d / 10
		sum += d % 10

		isSecond = !isSecond
	}
	return sum%10 == 0
}
