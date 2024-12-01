package helpers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nexivia/ms-lib-framework/internal/constants"
)

func CheckConnected() {
	if connected() {
		fmt.Println(constants.ConnectionSuccessful)
	} else {
		log.Fatal(constants.ConnectionFailed)
	}
}

func connected() bool {
	_, err := http.Get("http://clients3.google.com/generate_204")
	return err == nil
}
