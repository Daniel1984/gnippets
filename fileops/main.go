package main

import (
	"fmt"
	"log"
	"os"
)

//func main() {
//	// Open file and create scanner on top of it
//	file, err := os.Open("OMNIPROD.DAT")
//	if err != nil {
//		log.Fatal(err)
//	}
//	scanner := bufio.NewScanner(file)
//
//	// Default scanner is bufio.ScanLines. Lets use ScanWords.
//	// Could also use a custom function of SplitFunc type
//	scanner.Split(bufio.ScanLines)
//
//	for scanner.Scan() {
//		fmt.Println(scanner.Text())
//	}
//
//	err = scanner.Err()
//	if err == nil {
//		log.Println("Scan completed and reached EOF")
//	} else {
//		log.Fatal(err)
//	}
//}

func main() {
	file, err := os.Open("./OMNIPROD.DATS")
	if err != nil {
		log.Fatal(err.Error())
	}

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%+v", stat)
}
