package main

import (
	"fmt"
	quasar_fire_core "github.com/pankas87/quasar-fire-core"
	"log"
	"net/http"
	"os"
)

// TODO: Routing according to the project requirements
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}

	message := quasar_fire_core.GetMessage([]string{"asd", "def"}, []string{"asd", "def"})
	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting port %s", port)
	}

	log.Printf("Listening on port %s", port)

	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatal(err)
	}
}
