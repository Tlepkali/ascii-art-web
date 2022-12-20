package output

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"os"
)

const (
	hashOfStandard   string = "81802daa5d0000076c70ba728c2deb73"
	hashOfShadow     string = "7f3a1a3893bc6ac8fe24e10cb78d7ff4"
	hashOfThinkertoy string = "34328000fcd99e635adea921d7615961"
)

func ReadBanner(font string) ([]byte, error) {
	var banner []byte
	banner, err := os.ReadFile("./output/banners/" + font + ".txt")
	if err != nil {
		return banner, err
	}
	hashOfBanner := getMD5Hash(string(banner))
	if hashOfBanner != hashOfStandard && hashOfBanner != hashOfShadow && hashOfBanner != hashOfThinkertoy {
		err = errors.New("ERROR: wrong hash")
		return banner, err
	}
	return banner, nil
}

func CheckNewLines(s []string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			return false
		}
	}
	return true
}

func CheckTxt(s string) bool {
	for _, v := range s {
		if v < 10 || v > 127 {
			return true
		}
	}
	return false
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
