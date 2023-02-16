package config

import (
	"crypto/tls"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gitlab.com/reinhardjs/go-epp-rest/internal/constants"
)

type Config struct {
	Cert *tls.Certificate
}

func InitConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, errors.Wrap(err, "init config: godotenv load:")
	}

	cfg := &Config{}

	trustoreFileName := os.Getenv(constants.PAY_WEB_CC_TRUSTORE_FILENAME)
	keystoreFileName := os.Getenv(constants.PAY_WEB_CC_KEYSTORE_FILENAME)
	cert, err := tls.LoadX509KeyPair(trustoreFileName, keystoreFileName)

	if err != nil {
		return nil, errors.Wrap(err, "init config: load x509 key pair:")
	}

	cfg.Cert = &cert

	return cfg, nil
}
