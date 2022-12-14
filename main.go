package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	log.Println("HTTP request", r, string(dump), err)
	if r.TLS == nil{
		w.Write([]byte("nil TLS"))
		return
	}
	log.Println("HTTP TLS", r.TLS, string(r.TLS.TLSUnique))
	certs := r.TLS.PeerCertificates
	log.Println("HTTP CERTS", certs)
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/", defaultHandler)
	server := &http.Server{
		Addr: ":8080",
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequestClientCert,
		},
		Handler: http.HandlerFunc(defaultHandler),
	}

	server.ListenAndServeTLS("server.crt", "server.key")
	// http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
}
