package entities

import "time"

type EPPPoll struct {
	Registry     string    `gorm:"registry"`
	Datetime     time.Time `gorm:"datetime"`
	MessageId    string    `gorm:"msgid"`
	MessageCount string    `gorm:"msgcount"`
	Message      string    `gorm:"msg"`
	QDate        time.Time `gorm:"qdate"`
	Domain       string    `gorm:"domain"`
	Status       string    `gorm:"status"`
	RequestDate  time.Time `gorm:"redate"`
	ExpireDate   time.Time `gorm:"exdate"`
	ActionDate   time.Time `gorm:"acdate"`
	RequestId    string    `gorm:"reid"`
	ActionId     string    `gorm:"acid"`
}
