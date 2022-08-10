package text

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// remove all \n from a file and return the string
func RemoveReturns(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	var s string
	fr := bufio.NewReader(f)
	for {
		line, err := fr.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err != nil {
			break
		}
		s += line
	}
	return s, nil
}

// remove spaces from a string
func RemoveSpaces(str string) (string, error) {
	var ans string
	for _, v := range str {
		if v != ' ' {
			ans += string(v)
		}
	}
	return ans, nil
}

func RR(str string) error {
	s, err := RemoveReturns(str)
	if err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}
