package server

import (
	"net/http"
	"io"
	"github.com/sanjid133/mutual-tls/pkg/cert"
	"crypto/x509"
	"crypto/tls"
	"fmt"
)


func Init() error {
	http.HandleFunc("/hello", HelloServer)

	c, err := cert.New()
	if err != nil {
		return err
	}

	//generate a server cert
	if err := c.Server(); err != nil {
		fmt.Println(err)
		return err
	}

	caCert := c.Store.CACert()
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs: caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
		SessionTicketsDisabled:   true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		NextProtos: []string{"h2", "http/1.1"},
	}

	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr: ":8080",
		TLSConfig:tlsConfig,
	}
	fmt.Println("Running server")
	return server.ListenAndServeTLS(c.Store.CertFile("server"), c.Store.KeyFile("server"))

}

func HelloServer(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "hello, world!\n"+ req.TLS.PeerCertificates[0].Subject.CommonName)
}