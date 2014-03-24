package main

import (
	"net/http"
	"os"
)

var (
	listenAddr = getenvDefault("CONTENT_STORE_ADDR", ":8080")
)

func getenvDefault(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}

	return val
}

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
})

func main() {

	http.ListenAndServe(listenAddr, handler)
}