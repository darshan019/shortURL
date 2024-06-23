package models

type stored_url struct {
	id        int
	Short_url string `json:"short_url"`
	Long_url  string `json:"long_url"`
	User_id   int    `json:"user_id"`
}

func StoredUrl() *stored_url {
	return &stored_url{}
}
