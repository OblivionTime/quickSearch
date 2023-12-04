package quickSearch

type Category struct {
	ID   int    `gorm:"primaryKey;column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}
