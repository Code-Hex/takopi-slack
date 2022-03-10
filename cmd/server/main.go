package main

import (
	"net/http"

	p "github.com/Code-Hex/takopi-slack"
)

func main() {
	http.HandleFunc("/takopi", p.TakopiCommand)
	http.ListenAndServe(":8080", nil)
}
