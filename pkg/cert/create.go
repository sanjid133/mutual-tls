package cert

import (
	"os"
	"github.com/sanjid133/mutual-tls/util"
	"github.com/appscode/kutil/tools/certstore"
	"github.com/spf13/afero"
	"k8s.io/client-go/util/cert"
	"net"
)

type Cert struct {
	Store *certstore.CertStore
}

func New() (*Cert, error) {
	cm, err := certstore.NewCertStore(afero.NewOsFs(), util.PKIDir())
	return &Cert{Store:cm}, err
}
func (c *Cert) Server() error {
	if err := os.MkdirAll(util.BaseDir(), 0755); err != nil {
		return err
	}

	// create a self signed ca cert
	if err := c.Store.InitCA(); err != nil {
		return err
	}

	// issue a ca signed server cert
	sans := cert.AltNames{
		IPs: []net.IP{net.ParseIP("127.0.0.1")},
	}

	crt, key, err := c.Store.NewServerCertPair("127.0.0.1", sans)
	if err != nil {
		return err
	}
	if err = c.Store.WriteBytes("server", crt, key); err != nil {
		return err
	}
	return nil
}

func (c *Cert) Client() error  {
	// issue a ca signed client cert
	crt, key , err := c.Store.NewClientCertPair("cn-client")
	if err != nil {
		return err
	}

	if err = c.Store.WriteBytes("client", crt, key); err != nil {
		return err
	}
	return nil
}