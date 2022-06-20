package do_validator

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
