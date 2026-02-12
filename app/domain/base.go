package domain

import (
	"strings"

	"golang.org/x/text/unicode/norm"
)

func convertAddress2EM(str string) string {
	hyphen := "―"
	replacers := []*strings.Replacer{
		// charsReplacer('a', 'z', 'ａ'),
		charsReplacer('A', 'Z', 'Ａ'),
		strings.NewReplacer("ー", hyphen, "-", hyphen, "―", hyphen, "-", hyphen, "－", hyphen, "−", hyphen),
		strings.NewReplacer("0", "〇", "1", "一", "2", "二", "3", "三", "4", "四", "5", "五", "6", "六", "7", "七", "8", "八", "9", "九"),
		strings.NewReplacer("\n", "<br>"),
		strings.NewReplacer("<br>", "<br>　"),
	}
	str = norm.NFKC.String(str)
	for _, r := range replacers {
		str = r.Replace(str)
	}
	return str
}

func convertName2EM(name string) string {
	return strings.TrimSpace(strings.ReplaceAll(name, " ", "　"))
}

func charsReplacer(oldStart, oldEnd, newStart int) *strings.Replacer {
	chars := make([]string, 0)
	for c := oldStart; c <= oldEnd; c++ {
		chars = append(chars, string(rune(c)), string(rune(newStart+(c-oldStart))))
	}
	return strings.NewReplacer(chars...)
}
