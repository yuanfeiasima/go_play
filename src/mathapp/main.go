package main

import (
	"fmt"
	//"mymath"
	"github.com/drone/routes"
	"net/http"
	//"log"
	//"database/sql"
	//"github.com/go-sql-driver/mysql"
	"mathapp/db"
	//"database/sql"
	"encoding/asn1"
)

func getuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")

	if uid == "lulu" {
		fmt.Fprintf(w, "beautiful girl  %s", uid)
	} else {
		//fmt.Fprintf(w, "you are get user %s", uid)
		fmt.Fprintf(w, db.SelectNews())
	}

}

func insertNews(w http.ResponseWriter, r *http.Request){
	params := r.URL.Query()
	title := params.Get(":title")
	fmt.Println("title:%s", title)

}

func main() {
	//fmt.Println("Hello, world. Sqrt(2) = %v\n", mymath.Sqrt(2))
	fmt.Println("mathapp start!")
	mux := routes.New()
	mux.Get("/user/:uid", getuser)
	mux.Get("/news/:title", insertNews)
	http.Handle("/", mux)
	http.ListenAndServe(":8088", nil)
	//log.Fatal()


}
