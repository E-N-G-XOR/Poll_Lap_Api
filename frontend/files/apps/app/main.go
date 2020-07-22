package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Labs struct {
	Id          int
	Name        string
	TimeStamps  string
	UpdateStamp string
	CreateStamp string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_DB")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	//db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM apptable ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Labs{}
	res := []Labs{}
	for selDB.Next() {
		var id int
		var name, timestamps, updatestamp, createstamp string
		err = selDB.Scan(&id, &name, &timestamps, &updatestamp, &createstamp)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.TimeStamps = timestamps
		emp.UpdateStamp = updatestamp
		emp.CreateStamp = createstamp
		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		timestamps := r.FormValue("timestamps")
		insForm, err := db.Prepare("INSERT INTO apptable(name, timestamps) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, timestamps)
		log.Println("INSERT: Name: " + name + " | TimeStamps: " + timestamps)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:5000")
	http.HandleFunc("/", Index)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/new", New)
	http.ListenAndServe(":5000", nil)
}
