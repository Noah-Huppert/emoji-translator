package main

import (
	"bufio"
	"os"
	"fmt"

	"github.com/Noah-Huppert/golog"
)

// charMap between runs and emoji's
// First item in value array is the regional indicator symbol for the char, if a second value exists it is the blood type value
var charMap map[rune][]string = map[rune][]string{
	'a': []string{"🇦", "🅰️"},
	'b': []string{"🇧", "🅱️"},
	'c': []string{"🇨"},
	'd': []string{"🇩"},
	'e': []string{"🇪"},
	'f': []string{"🇫"},
	'g': []string{"🇬"},
	'h': []string{"🇭"},
	'i': []string{"🇮"},
	'j': []string{"🇯"},
	'k': []string{"🇰"},
	'l': []string{"🇱"},
	'm': []string{"🇲"},
	'n': []string{"🇳"},
	'o': []string{"🇴", "🅾️"},
	'p': []string{"🇵"},
	'q': []string{"🇶"},
	'r': []string{"🇷"},
	's': []string{"🇸"},
	't': []string{"🇹"},
	'u': []string{"🇺"},
	'v': []string{"🇻"},
	'w': []string{"🇼"},
	'x': []string{"🇽"},
	'y': []string{"🇾"},
	'z': []string{"🇿"},
}

func main() {
	// Setup logger
	logger := golog.NewStdLogger("emoji-translate")
	
	// Read stdin
	reader := bufio.NewReader(os.Stdin)
	
	txt, err := reader.ReadString('\n')
	if err != nil {
		logger.Fatalf("failed to read stdin")
	}

	// Translate
	out := ""
	
	for _, c := range txt {
		if emojis, ok := charMap[c]; ok {
			out += emojis[0]
		} else {
			out += string(c)
		}
	}

	fmt.Println(out)
}
