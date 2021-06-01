package main

import (
	"bytes"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Member struct {
	Name  string
	Email string
	Date  string
}

//go:embed static
var staticFiles embed.FS

//go:embed static/favicon.ico
var faviconFile []byte

//go:embed templates
var indexHTML embed.FS

var members []Member

func main() {
	t, err := template.ParseFS(indexHTML, "templates/*.html.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		name := req.Form.Get("member-name")
		email := req.Form.Get("member-email")
		if name != "" && email != "" {
			members = append(members, Member{
				Name:  name,
				Email: email,
				Date:  time.Now().Format("02.01.2006"),
			})
		}

		w.Header().Add("Content-Type", "text/html")

		t.ExecuteTemplate(w, "index.html.tmpl", struct {
			Members []Member
		}{Members: members})
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

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, "favicon.ico", time.Now(), bytes.NewReader(faviconFile))
	})
}
