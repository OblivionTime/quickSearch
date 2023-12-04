package quickSearch

type FileInfo struct {
	ID       int    `gorm:"primaryKey;column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	FileType string `gorm:"column:file_type;" json:"file_type"`
	ColumnID int    `gorm:"column:column_id" json:"column_id"`
	FilePath string `gorm:"column:file_path;type:longtext" json:"file_path"`
}
