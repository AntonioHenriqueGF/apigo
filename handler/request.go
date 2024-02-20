package handler

import "fmt"

func errParamIsRequired(name string, typ string) error {
	return fmt.Errorf("[param] missing required field '%s' (of type '%s')", name, typ)
}

func errEntrieNotFound(entrieName string) error {
	return fmt.Errorf("[query] %s not found", entrieName)
}

func errBadBodyFormating() error {
	return fmt.Errorf("[payload] body is empty or malformed")
}

type CreatePostRequest struct {
	Title   string `json:"pst_title" binding:"required"`
	Content string `json:"pst_content" binding:"required"`
}

func (r *CreatePostRequest) Validate() error {
	if *r == (CreatePostRequest{}) {
		return errBadBodyFormating()
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}

	if r.Content == "" {
		return errParamIsRequired("content", "string")
	}

	return nil
}
