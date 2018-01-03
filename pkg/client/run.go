package client

import (
	"crypto/tls"
	"github.com/sanjid133/mutual-tls/pkg/cert"
	"crypto/x509"
	"net/http"
	"fmt"
	"io/ioutil"
)

func Run() error  {
	c, err := cert.New()
	if err != nil {
		return err
	}
	if err = c.Store.LoadCA(); err != nil {
		return err
	}

	// create client cert
	if err = c.Client();err != nil {
		return err
	}

	//Load client cert
	crt, err := tls.LoadX509KeyPair(c.Store.CertFile("client"), c.Store.KeyFile("client"))
	if err != nil {
		return err
	}

	//serverCert, _, err := c.Store.ReadBytes("server")//CACert()

	// get the ca public cert
	caCert := c.Store.CACert()

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{crt},
		RootCAs: caCertPool,
	}

	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig:tlsConfig}
	client := &http.Client{Transport:transport}

	resp, err := client.Get("https://127.0.0.1:8080/hello")
	if err != nil {
		return err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(contents))
	return nil
}
