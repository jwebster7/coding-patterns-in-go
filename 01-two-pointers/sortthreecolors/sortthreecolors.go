package sortthreecolors

func sortThreeColors(colors []int) []int {
	if len(colors) == 0 {
		return colors
	}

	red, white, blue := 0, 0, len(colors)-1
	for white <= blue {
		switch colors[white] {
		case 0:
			if colors[red] != 0 {
				colors[white], colors[red] = colors[red], colors[white]
			}
			red++
			white++
		case 1:
			white++
		case 2:
			if colors[blue] != 2 {
				colors[white], colors[blue] = colors[blue], colors[white]
			}
			blue--
		}
	}

	return colors
}
