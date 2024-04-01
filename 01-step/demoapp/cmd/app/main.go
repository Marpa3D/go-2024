package main

import (
	"fmt"
	"net/http"

	"github.com/Marpa3D/go-2024/01-step/demoapp/pkg/stringreverse"
)

func main() {
	fmt.Println("Start server. Port 8080...")
	http.ListenAndServe(":8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s := r.URL.Query()["s"]
			rev := stringreverse.Revstr(s[0])
			w.Write([]byte(rev))

		}),
	)
}
