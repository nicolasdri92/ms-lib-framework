package framework

import (
	"fmt"
	"log"
	"net/http"
)

func CheckConnected() {
	if connected() {
		fmt.Println(ConnectionSuccessful)
	} else {
		log.Fatal(ConnectionFailed)
	}
}

func connected() bool {
	_, err := http.Get("http://clients3.google.com/generate_204")
	return err == nil
}
