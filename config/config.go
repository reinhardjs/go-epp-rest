package config

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/constants"
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure/mysql"
)

type Config struct {
	PayWebCCRootCaCert *x509.CertPool
	PayWebCCCert       *tls.Certificate
	Mysql              *mysql.Config
}

func InitConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.Wrap(err, "init config: godotenv load")
	}

	cfg := &Config{}

	mysqlHost := os.Getenv(constants.MYSQL_HOST)
	mysqlDBName := os.Getenv(constants.MYSQL_DB)
	mysqlUser := os.Getenv(constants.MYSQL_USER)
	mysqlPassword := os.Getenv(constants.MYSQL_PASSWORD)
	cfg.Mysql = &mysql.Config{
		Host:     mysqlHost,
		DBName:   mysqlDBName,
		User:     mysqlUser,
		Password: mysqlPassword,
	}

	trustoreFileName := os.Getenv(constants.PAY_WEB_CC_TRUSTORE_FILENAME)
	keystoreFileName := os.Getenv(constants.PAY_WEB_CC_KEYSTORE_FILENAME)
	payWebCCCert, err := tls.LoadX509KeyPair(trustoreFileName, keystoreFileName)
	if err != nil {
		return nil, errors.Wrap(err, "init config: load x509 key pair")
	}
	cfg.PayWebCCCert = &payWebCCCert

	caCert, err := ioutil.ReadFile(os.Getenv(constants.PAY_WEB_CC_CA_CERT_FILENAME))
	if err != nil {
		return nil, errors.Wrap(err, "init config: reading ca cert file")
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		return nil, errors.Wrap(err, "init config: append ca certs from PEM")
	}
	cfg.PayWebCCRootCaCert = caCertPool

	return cfg, nil
}
