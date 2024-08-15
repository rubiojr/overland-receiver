package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"flag"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	dir := flag.String("dir", "", "The directory to save files. Defaults to the current dir")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", savePayload(*dir))
	if err := http.ListenAndServe(":3111", r); err != nil {
		fmt.Println(err)
	}
}

func savePayload(dir string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fileName := fmt.Sprintf("%d/overland-%d.geojson", dir, time.Now().UnixNano())
		err = os.WriteFile(fileName, body, 0644)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result": "ok"}`))
	}
}
