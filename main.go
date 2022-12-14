package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	log.Println("HTTP request", r, string(dump), err)
	log.Println("HTTP TLS", r.TLS, string(r.TLS.TLSUnique))
	certs := r.TLS.PeerCertificates
	log.Println("HTTP CERTS", certs)
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
}
