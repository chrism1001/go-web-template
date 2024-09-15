package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"portfolio/internal/middleware"
	"portfolio/internal/template"

	"github.com/joho/godotenv"
)

func serveStaticfiles(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):]
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func main() {

	_ = godotenv.Load()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /static/", serveStaticfiles)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		middleware.Chain(w, r, template.Home("Christopher Mejia"))
	})

	fmt.Printf("server is running on port %s\n", os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	if err != nil {
		fmt.Println(err)
	}
}
