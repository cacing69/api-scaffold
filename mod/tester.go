package mod

import (
	_ "github.com/jinzhu/gorm"
)

type Tester struct {
	ID    uint   `gorm:"primary_key;column:tester_id"`
	Key   string `gorm:"primary_key;column:tester_key"`
	Value string `gorm:"primary_key;column:tester_value"`
}

func (m Tester) TableName() string {
	return "m_tester"
}
