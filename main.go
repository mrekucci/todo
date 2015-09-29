// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mrekucci/todo/internal/task"
)

var addrFlag = flag.String("addr", "127.0.0.1:8080", "address:port on which the server will be listening")

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: todo [options]\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	flag.PrintDefaults()
	os.Exit(2)
}

// corsHeaders wraps http.HandlerFunc and adds CORS headers to response.
// Function also handles pre-flight requests.
func corsHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" { // Stop the pre-flight request.
			return
		}
		fn(w, r)
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("todo: ")
	flag.Usage = usage
	flag.Parse()

	http.Handle(task.Path, http.HandlerFunc(corsHeaders(task.RestAPI)))
	http.Handle("/", http.FileServer(http.Dir("frontend/web")))
	if err := http.ListenAndServe(*addrFlag, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
