func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	a1 := area(ax1, ay1, ax2, ay2)
	a2 := area(bx1, by1, bx2, by2)
	if !intersect(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2) {
		return a1 + a2
	}

	x := []int{ax1, ax2, bx1, bx2}
	y := []int{ay1, ay2, by1, by2}

	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	sort.Slice(y, func(i, j int) bool {
		return y[i] < y[j]
	})

	in := area(x[1], y[1], x[2], y[2])

	return a1 + a2 - in
}

func area(x1, y1, x2, y2 int) int {
	return (x2 - x1) * (y2 - y1)
}

func intersect(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) bool {
	if ax1 <= bx1 && bx1 <= ax2 {
		return intersectY(ax1, ay1, ax2, ay2, bx1,by1,bx2,by2)
	}

	if ax1 <= bx2 && bx2 <= ax2 {
		return intersectY(ax1, ay1, ax2, ay2, bx1,by1,bx2,by2)
	}

	if bx1 <= ax1 && ax1 <= bx2 {
		return intersectY(ax1, ay1, ax2, ay2, bx1,by1,bx2,by2)
	}

	if bx1 <= ax2 && ax2 <= bx2 {
		return intersectY(ax1, ay1, ax2, ay2, bx1,by1,bx2,by2)
	}
	return false
}
func intersectY(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) bool {
	if ay1 <= by1 && by1 <= ay2 {
		return true
	}

	if ay1 <= by2 && by2 <= ay2 {
		return true
	}

	if by1 <= ay1 && ay1 <= by2 {
		return true
	}

	if by1 <= ay2 && ay2 <= by2 {
		return true
	}
	return false
}
