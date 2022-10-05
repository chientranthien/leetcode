
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	a1 := area(ax1, ay1, ax2, ay2)
	a2 := area(bx1, by1, bx2, by2)

	x1 := ax1
	if bx1 > x1 {
		x1 = bx1
	}

	x2 := ax2
	if bx2 < x2 {
		x2 = bx2
	}

	y1 := ay1
	if by1 > y1 {
		y1 = by1
	}

	y2 := ay2
	if by2 < y2 {
		y2 = by2
	}
	common := 0

	if x1 < x2 && y1 < y2 {
		common = area(x1, y1, x2, y2)
	}
	return a1 + a2 - common
}

func area(x1, y1, x2, y2 int) int {
	return (x2 - x1) * (y2 - y1)
}