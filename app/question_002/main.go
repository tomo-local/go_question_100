package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	var go_en string = "Go"
	var go_ja string = "Go言語"

	fmt.Printf("%s length:%d\n", go_en, len(go_en))
	fmt.Printf("%s length:%d\n", go_ja, len(go_ja))

	fmt.Printf("%s length:%d\n", go_en, utf8.RuneCountInString(go_en))
	fmt.Printf("%s length:%d\n", go_ja, utf8.RuneCountInString(go_ja))
}
