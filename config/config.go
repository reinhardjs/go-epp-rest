package config

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
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
	env := os.Getenv("ENV")

	if env != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, errors.Wrap(err, "init config: godotenv load")
		}
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

	trustore, err := base64.StdEncoding.DecodeString(os.Getenv(constants.PAY_WEB_CC_TRUSTORE))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode trustore base64")
	}

	keystore, err := base64.StdEncoding.DecodeString(os.Getenv(constants.PAY_WEB_CC_KEYSTORE))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode keystore base64")
	}

	cert, err := parseCertificate([]byte(trustore))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse certificate:")
	}

	key, err := parsePrivateKey([]byte(keystore))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse private key:")
	}

	tlsCert := tls.Certificate{
		Certificate: [][]byte{cert},
		PrivateKey:  key,
	}
	cfg.PayWebCCCert = &tlsCert

	// caCert, err := base64.StdEncoding.DecodeString(os.Getenv(constants.PAY_WEB_CC_CA_CERT))
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Failed to decode caCert base64")
	// }

	// caCertPool := x509.NewCertPool()
	// if !caCertPool.AppendCertsFromPEM(caCert) {
	// 	return nil, errors.Wrap(err, "init config: append ca certs from PEM")
	// }
	// cfg.PayWebCCRootCaCert = caCertPool

	return cfg, nil
}

func parseCertificate(pemBlock []byte) ([]byte, error) {
	block, _ := pem.Decode(pemBlock)
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return cert.Raw, nil
}

func parsePrivateKey(pemBlock []byte) (interface{}, error) {
	block, _ := pem.Decode(pemBlock)
	if block == nil {
		return nil, fmt.Errorf("failed to parse private key PEM")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
