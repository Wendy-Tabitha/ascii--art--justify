package graphics

// This function is used to show the graphics by calculating the value of each line and returning the string version of the graphic.
func Graphic(str, banner []string) string {
	var lines int
	result := ""
	// align := ""

	for _, ch := range str {
		if ch == "" {
			lines = 1
		} else {
			lines = 8
		}
		for i := 1; i <= lines; i++ {
			for _, runechar := range ch {
				index := (runechar-32)*9 + rune(i)
				result += banner[index]
			}
			// align += Align(result, alignType)
			result += "\n"
		}
	}
	return result
}
