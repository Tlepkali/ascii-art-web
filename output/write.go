package output

import "strings"

func FilterAndWrite(text []string, stargs string, symbs []string) string {
	output := ""
	if len(text) == 1 && text[0] == "" {
		return output
	}
	for i := 0; i < len(text); i++ {
		if len(text) == len(stargs)+1 && text[i] == "" {
			text = append(text[:i], text[i+1:]...)
		}
	}
	onlyNewLines := CheckNewLines(text)
	for i, word := range text {
		if onlyNewLines && i == len(text)-1 {
			break
		}
		if word == "" {
			output += "\n"
		} else {
			for line := 0; line < 8; line++ {
				for _, rune := range word {
					output += strings.Split(symbs[rune-32], "\n")[line]
				}
				output += "\n"
			}
		}
	}
	return output
}
