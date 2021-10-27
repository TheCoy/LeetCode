package golang


var (
	dx = []int{1,0,0,-1}
	dy = []int{0,1,-1,0}
)

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	oldColor := image[sr][sc]
	if image[sr][sc] != newColor {
		dfsFill(image, sr, sc, oldColor, newColor)
	}

	return image
}

func dfsFill(image [][]int, sr, sc int, oldColor, newColor int) {
	if image[sr][sc] != oldColor{
		return
	}
	image[sr][sc] = newColor
	for i := 0; i < 4; i++ {
		mx, my := sr + dx[i], sc + dy[i]
		if mx >= 0 && my >= 0 && mx < len(image) && my < len(image[0]) {
			dfsFill(image, mx, my, oldColor, newColor)
		}
	}
}
