package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cooldev95/MUltiSigTendermint/utils"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	w.Write([]byte(message))
}

func main() {
	ports, _ := utils.GetPorts()

	for _, p := range ports {
		fmt.Println("sahith don: ", p)

	}
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
