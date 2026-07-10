package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func health(w http.ResponseWriter, req *http.Request) {
	// TODO: retourner un json propre
	fmt.Fprintf(w, "ok\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func redirectToTls(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://IPAddr:443"+r.RequestURI, http.StatusMovedPermanently)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// curl -v -H "Host: tritounet.fr" http://localhost
	target := url.URL{Scheme: "https", Host: "tritounet.fr", Path: "/"}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			if debug {
				dump, err := httputil.DumpRequest(r.In, true)
				if err != nil {
					log.Fatal("Error while trying to dump the request : ", err.Error())
				}
				fmt.Printf("Dump IN : %s\n", dump)
			}
			r.SetURL(&target)
			r.SetXForwarded()
			if debug {
				dump, err := httputil.DumpRequestOut(r.Out, true)
				if err != nil {
					log.Fatal("Error while trying to dump the request : ", err.Error())
				}
				fmt.Printf("Dump OUT : %s\n", dump)
			}
		},
	}

	mux.HandleFunc("/health", health)
	mux.HandleFunc("/headers", headers)
	mux.Handle("/", proxy)

	fmt.Println("Server starting on port 80")

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
