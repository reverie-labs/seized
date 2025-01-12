package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func seizedHandler(w http.ResponseWriter, r *http.Request) {
	// parse to see if we've got an env var
	hardcoded_error := os.Getenv("SEIZED_HARDCODED_ERROR")

	ans := ""
	if hardcoded_error != "" {
		// if we have a hardcoded error, use it
		ans = hardcoded_error
	} else {
		// otherwise, use whatever we got on the request
		ans = r.URL.Path[1:]
	}

	// make an integer out of it
	ans_i, err := strconv.Atoi(ans)

	if err != nil {
		// if that didn't properly look like an integer, then 500 looks good enough
		ans_i = 500
	}

	// send the error code out immediately
	w.WriteHeader(ans_i)

	// read based on the integer not the string, rudimentary traversal protection
	fp := path.Join("static", strconv.Itoa(ans_i)+".html")
	file, err := os.Open(fp)

	// if that didn't work, then read the 500 error code
	if err != nil {
		// by this point it's too late to change the returned code, but I'm calling this a feature
		file, _ = os.Open("static/500.html")
	}

	// return the data
	http.ServeContent(w, r, "", time.Time{}, file)
}

func main() {
	log.Print("seized listening on :4000...")

	http.HandleFunc(`/`, seizedHandler)
	http.HandleFunc(`/favicon.ico`, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		// log.Fatal(err)
	}
}
