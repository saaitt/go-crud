package request

import "errors"

type CreateProductRequest struct {
	Title    string `json:"title"`
	Description string `json:"description"`
}

func (c CreateProductRequest) Validate() error {
	if c.Title == "" {
		return errors.New("title is required")
	}
	if c.Description == "" {
		return errors.New("description is required")
	}
	return nil
}