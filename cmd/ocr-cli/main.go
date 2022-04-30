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
	result := recognizer.Recognize(os.Args[1])
	if result.Error != nil {
		errMsg := result.Message
		if errMsg == "" {
			errMsg = fmt.Sprintf("%s", result.Error)
		}
		fmt.Printf("recognize failed: %s\n", errMsg)
	} else {
		fmt.Println(result.Image)
		for i, observation := range result.Observations {
			fmt.Println(i, observation)
		}
	}
}
