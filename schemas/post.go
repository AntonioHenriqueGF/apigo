package schemas

import "time"

type Post struct {
	ID            int        `json:"pst_id" db:"pst_id"`
	Title         string     `json:"pst_title" db:"pst_title" binding:"required"`
	Date_creation *time.Time `json:"pst_date_creation" db:"pst_date_creation" binding:"required"`
	Date_edit     *time.Time `json:"pst_date_edit" db:"pst_date_edit" binding:"required"`
	Content       string     `json:"pst_content" db:"pst_content" binding:"required"`
}
