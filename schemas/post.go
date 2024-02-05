package schemas

type Post struct {
	ID            int    `json:"pst_id"`
	Title         string `json:"pst_title"`
	Date_creation string `json:"pst_date_creation"`
	Date_edit     string `json:"pst_date_edit"`
	Content       string `json:"pst_content"`
}
