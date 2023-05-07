package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", func(w http.ResponseWriter, req *http.Request) {
		_, err := fmt.Fprint(w, "HELLO WORLD 333")
		if err != nil {
			return
		}
	})

	fmt.Println("AAAA3")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
