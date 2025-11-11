package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fullsailor/pkcs7"
	"howett.net/plist"
)

func checkProvision(directory string) bool {
	provisionFile := directory

	data, err := os.ReadFile(provisionFile)
	if err != nil {
		log.Fatalln("failed to read file: ", err)
	}

	p7, err := pkcs7.Parse(data)
	if err != nil {
		log.Fatalln("failed to parse cms data: ", err)
	}

	var provision map[string]interface{}
	_, err = plist.Unmarshal(p7.Content, &provision)
	if err != nil {
		log.Fatalln("failed to parse plist: ", err)
	}

	if _, ok := provision["PPQCheck"]; ok {
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

	if ppqCheck {
		println("PPQ has been detected!")
		return
	} else {
		println("PPQ has not been detected!")
		return
	}
}
