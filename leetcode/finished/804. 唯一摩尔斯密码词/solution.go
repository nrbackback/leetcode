package main

func uniqueMorseRepresentations(words []string) int {
	morseTable := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-",
		".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	existsCode := make(map[string]bool)
	for _, v := range words {
		morseCode := ""
		for _, v1 := range v {
			morseCode += morseTable[v1-'a']
		}
		existsCode[morseCode] = true
	}
	return len(existsCode)
}
