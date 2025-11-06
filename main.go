package main

import (
	"log"
	"os"
	"regexp"
)

func main() {
	provisionFile := "345345355.mobileprovision"

	data, err := os.ReadFile(provisionFile)
	if err != nil {
		log.Fatal(err)
	}

	pattern := `<key>PPQCheck</key>\s*<(true|false)/?`

	re := regexp.MustCompile(pattern)
	match := re.FindString(string(data))

	if match != "" {
		println("ppqcheck")
	} else {
		println("no ppqcheck")
	}
}
