package main

import "testing"

func TestReplaceKanSuji(t *testing.T) {
	res := replaceKanSuji("大阪府箕面市桜ケ丘4-3-127")
	t.Error(res)
}

func TestShowReplacerAlphabet(t *testing.T) {
	// t.Error(replacerAlphabet)
}
