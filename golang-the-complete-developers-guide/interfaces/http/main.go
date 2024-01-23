package main

import (
	"fmt"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	bs := make([]byte, 9999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}

func (logWriter) Write(bs []byte) (int, error) {
	return 1, nil
}
