package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func listdir(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path == "" {
		path = "./"
	}
	if strings.HasSuffix(path, "/") {
		fmt.Println("LIST ", path)
		files, _ := ioutil.ReadDir(path)

		fmt.Fprintf(w, "<html><body><h1>%s</h1><Hr />", path)

		if path != "./" {
			fmt.Fprintf(w, "<a href='..' >..<Br />")
		}

		for _, f := range files {
			isdir := ""
			if f.IsDir() && ! strings.HasSuffix(f.Name(), "/") {
				isdir = "/"
			}
			fmt.Fprintf(w, "<a href='/%s/%s%s' >%s<Br />", path, f.Name(), isdir, f.Name())
		}
	} else {
		fmt.Println("READ ", path)
		body, _ := ioutil.ReadFile(path)
		fmt.Fprintf(w, string(body))
	}

	
}

func main() {
	fmt.Println("Listening on 0.0.0.0:8080")
	http.HandleFunc("/", listdir)
	http.ListenAndServe(":8080", nil)
}
