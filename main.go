package main

import (
	"fmt"
	"net/http"
	"os"
	"portfolio/internal/middleware"
	"portfolio/internal/template"
	"portfolio/internal/view"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /static/", view.ServeStaticFiles)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		middleware.Chain(w, r, template.Home())
	})

	fmt.Printf("server is running on port %s\n", os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	if err != nil {
		fmt.Println(err)
	}
}
