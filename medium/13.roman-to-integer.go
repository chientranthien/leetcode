
func romanToInt(s string) int {
	b2i := map[byte]int{
		'I' :1,
		'V' :5,
		'X' :10,
		'L' :50,
		'C' :100,
		'D' :500,
		'M' :1000,
	}

	b := []byte(s)
	sum := 0
	i := 0

	for ; i<len(s)-1; i++ {
			v1 := b2i[b[i]]
			v2 := b2i[b[i+1]]
			if v1 < v2 {
				sum += v2-v1
				i++
			} else{
				sum += v1
			}
	}

	if i < len(b) {
			sum += b2i[b[i]]
	}

	return sum
}
