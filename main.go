package main

import (
	"fmt"
	"github.com/betology/romanNumerals"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	// http package has methods for dealing with request
	http.HandleFunc("/", func(w http.ResposeWriter, r *http.Request) {
		urlPathElements := string.Split(r.URL.Path, "/")
		// If request is GET with correct syntax
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(string.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				// If resource is not in teh list, send Not Found status
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				fmt.Fprintf(w, "q", html.EscapeString(romanNumerals.Numerals[number]))
			}
		} else {
			// for all other request, tell that Cleent sent a bad request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad reqeust"))
		}
	})
	// Create a server and run it on 8000 port
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe
}
