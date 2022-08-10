package text

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"
)

type RegexSplitter struct {
	Splitter *regexp.Regexp
}

func NewRegexSplitter(pattern string) *RegexSplitter {
	return &RegexSplitter{
		Splitter: regexp.MustCompile(pattern),
	}
}

func (s *RegexSplitter) Split(text string) []string {
	return s.Splitter.Split(text, -1)
}

// reverse string slice in O(n) time
func Reverse(slice []string) []string {
	ans := make([]string, len(slice))
	for i := 0; i < len(slice); i++ {
		ans[i] = slice[len(slice)-1-i]
	}
	return ans
}

func Srv(i, o, pattern string) error {
	if i == "i" {
		i = def_i
	}
	if pattern == "pcg" {
		pattern = pcg
	}
	if pattern == "pcg-" {
		pattern = pcg_
	}
	if o == "o" {
		o = def_o
	}
	// read string from file i
	f, err := os.Open(i)
	if err != nil {
		return err
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	// []byte to string
	i = string(content)
	if err != nil {
		return err
	}
	ans := NewRegexSplitter(pattern).Split(i)
	for _, v := range Reverse(ans) {
		println(v)
	}
	ans = Reverse(ans)
	// write ans to file output, if exists, overwrite it
	// _, err = os.Stat(o)
	// if err == nil {
	// 	os.Remove(o)
	// }
	file, err := os.Create(o)
	if err != nil {
		return errors.New("create file error " + err.Error() + " at:" + o)
	}
	defer file.Close()
	for _, v := range ans {
		if v == "" {
			continue
		}
		file.WriteString(v + ",")
	}
	return nil
}

var (
	pcg   = `[+-]\d+(?:\.\d+)?%`
	pcg_  = `[+-]\d+(?:\.\d+)?%|[-]`
	def_i = `C:\test\i.txt`
	def_o = `C:\test\o.csv`
)
