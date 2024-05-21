package other

import "fmt"

var a rune = 1072
var ya rune = 1103
var space rune = 32

var encryptedMsg string = "Ебгбк рсйефнбён есфдпк щйхс"

func Decrypt() {
	for i := 0; i < 1103-1072; i++ {
		var decryptedMsg = ""
		for _, val := range []rune(encryptedMsg) {
			if val == space {
				decryptedMsg += " "
			} else {
				decryptedMsg += string(val - rune(i))
			}
		}
		fmt.Println(decryptedMsg)
	}
}
