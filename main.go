package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"crypto-server/handler"
)

func main() {
	port := 8080

	if strValue, ok := os.LookupEnv("PORT"); ok {
		if intValue, err := strconv.Atoi(strValue); err == nil {
			port = intValue
		}
	}

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		handler.NewHandler()),
	)
}
