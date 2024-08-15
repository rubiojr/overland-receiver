package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"flag"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	dir := flag.String("dir", "", "The directory to save files. Defaults to the current dir")
	address := flag.String("listen", ":3111", "Service listen address")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", savePayload(*dir))
	if err := http.ListenAndServe(*address, r); err != nil {
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

		today := time.Now().Format("2006/01/2")
		destDir := filepath.Join(dir, today)
		os.MkdirAll(destDir, 0755)
		fileName := fmt.Sprintf("overland-%d.geojson", time.Now().UnixNano())
		dest := filepath.Join(destDir, fileName)
		err = os.WriteFile(dest, body, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result": "ok"}`))
		fmt.Println("Saved file:", fileName)
	}
}
