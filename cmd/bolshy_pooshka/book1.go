package main

import "math/rand"

//go:generate go-bindata -pkg $GOPACKAGE -o book1txt.go book1.txt

var txt = string(MustAsset("book1.txt"))

func GetRandText(maxLen int) string {
	txtLen := rand.Intn(maxLen/10*9) + maxLen/10
	txtPos := rand.Intn(len(txt) - txtLen)
	return txt[txtPos : txtPos+txtLen]
}
