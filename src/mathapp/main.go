package main

import (
	"fmt"
	//"mymath"
	"net/http"
	"github.com/drone/routes"
)

func getuser(w http.ResponseWriter, r *http.Request){
	params := r.URL.Query()
	uid := params.Get(":uid")

	fmt.Fprintln(w, "you are get user %s", uid)
}
func main() {
	//fmt.Println("Hello, world. Sqrt(2) = %v\n", mymath.Sqrt(2))
	mux := routes.New()
	mux.Get("/user/:uid", getuser)
	http.Handle("/", mux)
	http.ListenAndServe(":8088", nil)
}
