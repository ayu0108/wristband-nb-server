package main

import (
	"fmt"
	"log"
	"nb_server/mysql"
	"nb_server/tcpserver"
	"net/http"
	"text/template"
)

const (
	addr = ""
	port = ":8080"
)

// Webservice : open service to provide web view
func Webservice() {
	db, err := mysql.Initdb()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(db)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", ServeHTTP)
	fmt.Printf("Web service listening on %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// ReciveData : 接收到的資料
type ReciveData struct {
	id         int
	name       string
	addr       string
	mac        string
	data       string
	date       string
	connStatus string
}

// Recive : 接收資料結構
type Recive struct {
	Title       string
	ReciveDatas []ReciveData
}

// ServeHTTP : Deal with html files
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/dashboard.html"))
	if r.URL.Path != "/" { //路徑不是根目錄
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	data := LoadData()
	tmpl.Execute(w, data)
}

// LoadData : Loading recive data
func LoadData() []ReciveData {
	db, err := mysql.Initdb()

	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query(`SELECT * FROM recive order by date desc`) // check err
	defer rows.Close()

	var data []ReciveData
	for rows.Next() {
		var u ReciveData
		if err := rows.Scan(&u.id, &u.name, &u.addr, &u.mac, &u.data, &u.date, &u.connStatus); err != nil { // check err
			fmt.Println(err)
		}
		data = append(data, u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func main() {
	Webservice()
	go tcpserver.ServerTCP()
}
