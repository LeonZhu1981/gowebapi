package main

import (
	"fmt"
	"gowebapi/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		fmt.Fprintf(w, "Hello %s\r\n", "Leon.X.Zhu")
	})
	router.GET("/points/:user/:pointsysno", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		fmt.Fprintf(w, "You request current page, user name:%s, point sys no:%s", params["user"], params["pointsysno"])
	})
	router.GET("/points/:user", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		r.ParseForm()
		fmt.Fprintf(w, "You request points page, user name:%s, query string:%s", params["user"], r.Form)
	})
	router.ServeFiles("/public/*filepath", http.Dir("./public/"))

	err := http.ListenAndServe(":9090", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
