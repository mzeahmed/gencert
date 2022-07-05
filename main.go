package main

import (
	"fmt"
	"os"

	"training.go/gencert/cert"
	"training.go/gencert/pdf"
)

func main() {
	c, err := cert.New("Golang programming", "Bob Dylan", "2018-06-21")
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}

	var saver cert.Saver
	saver, err = pdf.New("output")
	if err != nil {
		fmt.Printf("Error during pdf generator creation: %v", err)
		os.Exit(1)
	}
	saver.Save(*c)
}
