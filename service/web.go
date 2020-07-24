package service

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	v1 "wristband-nb-server/api/v1"
)

const (
	// WebPort - web listen on port
	WebPort = ":8080"
)

// Webservice : open service to provide web view
func Webservice() {
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	http.HandleFunc("/", ServeHTTP)
	fmt.Printf("Web service listening on %s\n", WebPort)
	if err := http.ListenAndServe(WebPort, nil); err != nil {
		log.Fatal(err)
	}
}

// ServeHTTP : Deal with html files
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	if r.URL.Path != "/" { //路徑不是根目錄
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	var data = v1.GetReceiveDatas()
	tmpl.Execute(w, data)
}
