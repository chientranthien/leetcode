func intToRoman(num int) string {
	r := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	v := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	rs := ""
	i := len(r) - 1
	for num > 0 {
		tmp := num / v[i]
		if tmp > 0 {
			for j := 0; j < tmp; j++ {
				rs += r[i]
			}
			num = num % v[i]
		} else {
			i--
		}
	}

	return rs
}
