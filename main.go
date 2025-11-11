package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func checkProvision(directory string) bool {
	provisionFile := directory

	data, err := os.ReadFile(provisionFile)
	if err != nil {
		log.Fatal(err)
	}

	key := `<key>PPQCheck</key>`

	re := regexp.MustCompile(key)
	match := re.FindString(string(data))

	if match != "" {
		return true
	} else {
		return false
	}
}

func main() {
	directory := flag.String("directory", "", "Directory of the .mobileprovision (required)")
	flag.Parse()

	if *directory == "" {
		fmt.Println("error: --directory flag is needed")
		flag.Usage()
		os.Exit(1)
	}

	ppqCheck := checkProvision(*directory)

	if ppqCheck == true {
		println("PPQ has been detected!")
		return
	} else {
		println("PPQ has not been detected!")
		return
	}
}
