package output

import (
	"errors"
	"fmt"
	"strings"
)

func Asciiart(arg string, font string) (string, error) {
	if arg == "" || font == "" {
		err := errors.New("Incorrect argument")
		return "", err
	}
	banner, err := ReadBanner(font)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if CheckTxt(arg) {
		err = errors.New("Incorrect argument")
		return "", err
	}
	symbs := strings.Split(strings.ReplaceAll(string(banner), "\r", ""), "\n\n")
	stargs := strings.ReplaceAll(arg, "\r\n", "\n")
	text := strings.Split(stargs, "\n")
	return FilterAndWrite(text, stargs, symbs), nil
}
