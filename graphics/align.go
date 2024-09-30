package graphics

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Align(alignment, artLine string) string {
	width := ScreenWidth()
	length := len(artLine)
	space := width - length

	fmt.Println(length)

	switch alignment {
	case "left":
		return artLine
	case "right":
		rightPadding := strings.Repeat(" ", space) + artLine
		return rightPadding
	}
	return ""
}

func ScreenWidth() int {
	cmd := exec.Command("stty", "size")

	cmd.Stdin = os.Stdin

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	str := strings.Fields(string(output))

	width, err := strconv.Atoi(str[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return width
}
