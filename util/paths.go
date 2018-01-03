package util

import (
	"path/filepath"
	"k8s.io/client-go/util/homedir"
)

func BaseDir() string  {
	return filepath.Join(homedir.HomeDir(), ".hoque", "ssl")
}

func PKIDir() string  {
	return  filepath.Join(BaseDir(), "pki")
}
