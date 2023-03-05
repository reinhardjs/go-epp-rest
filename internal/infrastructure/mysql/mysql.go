package mysql

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	DBName   string
	User     string
	Password string
}

func NewMysqlConnection(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.DBName)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.Wrap(err, "NewMysqlConnection: gorm.Open()")
	}

	return conn, nil
}
