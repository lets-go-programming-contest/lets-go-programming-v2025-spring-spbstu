package manager

type Contact struct {
	ID    string `gorm:"primaryKey;type:uuid"`
	Name  string `gorm:"not null"`
	Phone string `gorm:"not null;unique"`
}

func (Contact) TableName() string {
	return "contacts"
}
