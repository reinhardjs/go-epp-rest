package config

import (
	"crypto/tls"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/constants"
	"gitlab.com/merekmu/go-epp-rest/pkg/mysql"
)

type Config struct {
	PayWebCCCert *tls.Certificate
	Mysql        *mysql.Config
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

	return cfg, nil
}
