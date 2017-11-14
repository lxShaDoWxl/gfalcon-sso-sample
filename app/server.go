package main

import (
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/m0cchi/gfalcon/model"
	"net/http"
	"os"
	"strconv"
)

var db *sqlx.DB

func writeErrorPage(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, "%v\n", err)
	fmt.Fprintf(w, "SignIn Page --> https://saas.m0cchi.net/\n")
	fmt.Fprintf(w, "Team/ID/Password: gfalcon/gfadmin/secret")
}

func handle(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("gfalcon.session")
	if err != nil {
		writeErrorPage(w, err)
		return
	}
	userIID, err := r.Cookie("gfalcon.iid")
	if err != nil {
		writeErrorPage(w, err)
		return
	}
	IID, err := strconv.ParseUint(userIID.Value, 10, 32)
	if err != nil {
		writeErrorPage(w, err)
		return
	}

	session, err := model.GetSession(db, uint32(IID), sessionID.Value)
	if err != nil {
		writeErrorPage(w, err)
		return
	}

	if err = session.Validate(); err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	fmt.Fprintf(w, "Authenticated")
}

func main() {
	var port int
	var dbhost string
	var err error
	flag.IntVar(&port, "port", 8080, "service's port")
	flag.StringVar(&dbhost, "dbhost", "", "gfalcon's DB")
	flag.Parse()

	if dbhost == "" {
		fmt.Println("required --dbhost [host]")
		os.Exit(1)
	}

	db, err = sqlx.Connect("mysql", dbhost)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	http.HandleFunc("/", handle)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
