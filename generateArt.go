package main

import "strings"

func GenerateArt(file []byte, text string) string {
	fileString := string(file)
	fileString = strings.Replace(fileString, "\r", "", -1)
	lines := strings.Split(fileString, "\n")
	text_lines := strings.Split(text, "\\n")

	var output string

	for l := 0; l < len(text_lines); l++ {
		runes := []rune(text_lines[l])

		for i := 0; i < 8; i++ {
			for j := 0; j < len(runes); j++ {
				pos := int(runes[j] - 32)
				pos2 := 1 + (int(runes[j])-32)*9 + i

				letterLen := length(lines, pos)

				output += lines[pos2]
				output += spaces(letterLen - len(lines[pos2]))
			}
			output += "\n"
		}
	}
	return output
}

func spaces(n int) string {
	var out string
	for i := 0; i < n; i++ {
		out += " "
	}
	return out
}

func length(list []string, ind int) int {
	i := 1 + 9*ind

	max := 1
	for ; i < 1+9*ind+9; i++ {
		lineLength := len(list[i])
		if lineLength > max {
			max = lineLength
		}
	}
	return max
}
