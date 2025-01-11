package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func seizedHandler(w http.ResponseWriter, r *http.Request) {
	// parse the number out of the path
	ans := r.URL.Path[1:]
	ans_i, _ := strconv.Atoi(ans)

	// read based on the number, rudimentary traversal protection
	// file, _ := os.Open("static/" + strconv.Itoa(ans_i) + ".html")

	// read the hardcoded file
	file, _ := os.Open("static/451.html")

	// write the error code
	w.WriteHeader(ans_i)

	// return the data
	http.ServeContent(w, r, "", time.Time{}, file)
}

func main() {
	log.Print("seized listening on :4000...")

	http.HandleFunc(`/`, seizedHandler)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		// log.Fatal(err)
	}
}