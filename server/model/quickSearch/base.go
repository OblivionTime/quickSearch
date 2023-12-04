package quickSearch

type SearchResult struct {
	FileList     []Info     `json:"fileList"`
	CategoryList []Category `json:"categoryList"`
}
type Info struct {
	CategoryName string `json:"category_name"`
	FileInfo
}
