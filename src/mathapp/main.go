package main

import (
	"fmt"
	//"mymath"
	"github.com/drone/routes"
	"net/http"
)

func getuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")

	if uid == "lulu" {
		fmt.Fprintf(w, "beautiful girl  %s", uid)
	} else {
		fmt.Fprintf(w, "you are get user %s", uid)
	}

}
func main() {
	//fmt.Println("Hello, world. Sqrt(2) = %v\n", mymath.Sqrt(2))
	fmt.Println("mathapp start!")
	mux := routes.New()
	mux.Get("/user/:uid", getuser)
	http.Handle("/", mux)
	http.ListenAndServe(":8088", nil)
}
