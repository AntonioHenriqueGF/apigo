package schemas

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Date   string `json:"date"`
	Body   string `json:"body"`
}
