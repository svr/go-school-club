package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"text/template"
)

//go:embed static
var staticFiles embed.FS

//go:embed templates
var indexHTML embed.FS

func main() {
	t, err := template.ParseFS(indexHTML, "templates/index.html.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var path = req.URL.Path
		w.Header().Add("Content-Type", "text/html")

		t.Execute(w, struct {
			Title    string
			Response string
		}{Title: "test", Response: path})

	})

	serveStaticAssets()

	port := os.Getenv("GO_SCHOOL_CLUB_PORT")
	if port == "" {
		port = "3000"
	}

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveStaticAssets() {
	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))
}
