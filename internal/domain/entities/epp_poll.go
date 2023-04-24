package entities

import (
	"time"
)

type EPPPoll struct {
	Registry       string    `gorm:"column:registry"`
	Datetime       time.Time `gorm:"column:datetime"`
	MessageId      string    `gorm:"column:msgid"`
	MessageCount   string    `gorm:"column:msgcount"`
	Message        string    `gorm:"column:msg"`
	QDate          time.Time `gorm:"column:qdate"`
	Domain         string    `gorm:"column:domain"`
	Status         string    `gorm:"column:status"`
	RequestingDate time.Time `gorm:"column:redate"`
	ExpireDate     time.Time `gorm:"column:exdate"`
	ActingDate     time.Time `gorm:"column:acdate"`
	RequestingId   string    `gorm:"column:reid"`
	ActingId       string    `gorm:"column:acid"`
}

func (EPPPoll) TableName() string {
	return "tbl_epp_poll"
}
