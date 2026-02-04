package handler

import (
	"fmt"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func errBadParamFormating(name string, typ string) error {
	return fmt.Errorf("[param] field '%s' expected to be formatted as '%s'", name, typ)
}

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

type CreateUserRequest struct {
	Name     string `json:"usu_name" binding:"required"`
	Email    string `json:"usu_email" binding:"required"`
	Password string `json:"usu_password" binding:"required"`
}

func (r *CreateUserRequest) Validate() error {
	if *r == (CreateUserRequest{}) {
		return errBadBodyFormating()
	}
	if r.Name == "" {
		return errParamIsRequired("usu_name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("usu_email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("usu_password", "string")
	}
	if !emailRegex.MatchString(r.Email) {
		return errBadParamFormating("usu_email", "email")
	}
	return nil
}

type LoginUserRequest struct {
	Email    string `json:"usu_email" binding:"required"`
	Password string `json:"usu_password" binding:"required"`
}

func (r *LoginUserRequest) Validate() error {
	if *r == (LoginUserRequest{}) {
		return errBadBodyFormating()
	}
	if r.Email == "" {
		return errParamIsRequired("usu_email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("usu_password", "string")
	}
	if !emailRegex.MatchString(r.Email) {
		return errBadParamFormating("usu_email", "email")
	}
	return nil
}
