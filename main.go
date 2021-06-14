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

type TemplateData struct {
	Members []Member
	Member  Member
}

//go:embed static
var staticFiles embed.FS

//go:embed static/favicon.ico
var faviconFile []byte

//go:embed templates
var indexHTML embed.FS

var members []Member

func main() {
	t := template.Must(template.ParseFS(indexHTML, "templates/*.html.tmpl"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		templateData := TemplateData{Members: members, Member: Member{Name: "", Email: "", Date: ""}}

		if r.Method == http.MethodPost {
			r.ParseForm()

			member := &Member{
				Name:  r.Form.Get("member-name"),
				Email: r.Form.Get("member-email"),
				Date:  time.Now().Format("02.01.2006"),
			}

			if member.Validate(members) {
				members = append(members, *member)
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				templateData.Member = *member
			}
		}

		t.ExecuteTemplate(w, "index.html.tmpl", templateData)
	})

	serveStaticAssets()

	port := os.Getenv("GO_SCHOOL_CLUB_PORT")
	if port == "" {
		port = "3000"
	}

	err := http.ListenAndServe(":"+port, nil)
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
