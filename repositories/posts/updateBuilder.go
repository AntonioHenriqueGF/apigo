package posts

type UpdateBuilder struct {
	setParts []string
	args     []any
}

func (u *UpdateBuilder) Add(field string, value any) {
	u.setParts = append(u.setParts, field+" = ?")
	u.args = append(u.args, value)
}
