package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"

	"github.com/Noah-Huppert/golog"
)

const blankSpace string = "󠀠"

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
	'0': []string{"0️⃣"},
	'1': []string{"1️⃣"},
	'2': []string{"2️⃣"},
	'3': []string{"3️⃣"},
	'4': []string{"4️⃣"},
	'5': []string{"5️⃣"},
	'6': []string{"6️⃣"},
	'7': []string{"7️⃣"},
	'8': []string{"8️⃣"},
	'9': []string{"9️⃣"},
	'-': []string{"〰️"},
}

func main() {
	// Setup logger
	logger := golog.NewStdLogger("emoji-translate")
	
	// Read stdin
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		logger.Fatalf("failed to read stdin: %s", err.Error())
	}
	
	txt := strings.ToLower(string(b))

	// Translate
	out := ""
	
	for _, c := range txt {
		if emojis, ok := charMap[c]; ok {
			out += emojis[0] + blankSpace
		} else {
			out += string(c)
		}
	}

	fmt.Println(out)
}
