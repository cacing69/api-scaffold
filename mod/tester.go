package mod

type Tester struct {
	ID    uint   `json:"id" gorm:"primary_key;column:tester_id"`
	Key   string `json:"key" gorm:"column:tester_key"`
	Value string `json:"value" gorm:"column:tester_value"`
}

func (m Tester) TableName() string {
	return "m_tester"
}
