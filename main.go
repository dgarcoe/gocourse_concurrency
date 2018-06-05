package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const web1 string = "http://ipv4.download.thinkbroadband.com/20MB.zip"
const web2 string = "http://ipv4.download.thinkbroadband.com/5MB.zip"

func downloadFile(url string, path string, msg chan string) {

	out, err := os.Create(path)
	if err != nil {
		log.Printf("Error creating file: %s", err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error getting web: %s", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Printf("Error copying file: %s", err)
	}

	s := fmt.Sprintf("%s", url)

	msg <- s

}

func main() {

	msg := make(chan string, 1)
	go downloadFile(web1, "./file20MB", msg)
	go downloadFile(web2, "./file5MB_1", msg)
	go downloadFile(web2, "./file5MB_2", msg)

	for i := 0; i < 3; i++ {
		log.Printf("%s", <-msg)
	}
}
