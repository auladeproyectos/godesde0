package webserve

import "net/http"

func MiWebserver() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserve/index.html")
}
