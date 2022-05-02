package main

import (
	"fmt"
	"github.com/chamzzzzzz/ocr/bills"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: bills-ocr-cli file")
		os.Exit(1)
	}

	recognizer := &bills.Recognizer{}
	bill, err := recognizer.Recognize(os.Args[1])
	if err != nil {
		fmt.Println("recognize failed.")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("recognize successed.")
	fmt.Println(bill)
}
