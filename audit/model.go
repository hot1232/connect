package audit

import "time"

type Log struct {
	Id int32 `gorm:"type int(32);auto increment;primary key"`
	Uid int32 `gorm:"type int(32);not null"`
	Username string `gorm:"type varchar(256);not null"`
	Cmd string `gorm:"type varchar(4096);not null"`
	Time time.Time `gorm:"type datetime;DEFAULT CURRENT_TIMESTAMP"`
}

func (self *Log) TableName() string{
	return "audit_log"
}
