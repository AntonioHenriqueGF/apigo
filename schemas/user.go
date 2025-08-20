package schemas

import "time"

type User struct {
	ID         int        `json:"usu_id" db:"usu_id"`
	Name       string     `json:"usu_name" db:"usu_name"`
	Email      string     `json:"usu_email" db:"usu_email"`
	Created_at *time.Time `json:"usu_created_at" db:"usu_created_at"`
}
