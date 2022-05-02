package main

import (
	"fmt"
	"github.com/chamzzzzzz/ocr"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ocr file")
		os.Exit(1)
	}

	recognizer := &ocr.MacRecognizer{}
	result, err := recognizer.Recognize(os.Args[1])
	if err != nil {
		fmt.Println("recognize failed.")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("recognize successed.")
	fmt.Println(result.Image)
	for i, observation := range result.Observations {
		fmt.Println(i, observation)
	}
}
