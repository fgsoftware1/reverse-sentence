package reverselib

import(
	"strings"
)

func Reverse(sentence string) string {
	words := strings.Split(sentence, " ")
	reversedWords := make([]string, len(words))
	for i, j := 0, len(words)-1; i < len(words); i, j = i+1, j-1 {
		reversedWords[i] = words[j]
	}
	return strings.Join(reversedWords, " ")
}